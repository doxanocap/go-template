package service

import (
	"app/internal/manager/interfaces"
	"app/internal/model"
	"context"
	"encoding/json"

	"github.com/doxanocap/pkg/errs"
	"github.com/golang-jwt/jwt/v5"
	"github.com/spf13/viper"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	manager interfaces.IManager
}

func InitUserService(manager interfaces.IManager) *UserService {
	return &UserService{
		manager: manager,
	}
}

func (us *UserService) Create(ctx context.Context, body model.SignUp) (result *model.AuthResponse, err error) {
	found, err := us.manager.Repository().User().FindByEmail(ctx, body.Email)
	if err != nil {
		return
	}
	if found != nil {
		return nil, model.ErrUserAlreadyExist
	}
	salt, _ := bcrypt.GenerateFromPassword([]byte(""), 10)

	passwordWithSalt := append(body.Password, salt...)
	hashedPassword, err := bcrypt.GenerateFromPassword(passwordWithSalt, bcrypt.DefaultCost)
	if err != nil {
		return nil, errs.Wrap("generate password hash", err)
	}
	body.Password = hashedPassword

	user, err := us.manager.Repository().User().Create(ctx, body)
	if err != nil {
		return
	}

	userDTO := model.ParseUserDTO(*user)
	tokens, err := us.manager.Service().Auth().NewPairTokens(userDTO)

	result = &model.AuthResponse{
		User:   userDTO,
		Tokens: *tokens,
	}
	return
}

func (us *UserService) Authenticate(ctx context.Context, body model.SignIn) (result *model.AuthResponse, err error) {
	user, err := us.manager.Repository().User().FindByEmail(ctx, body.Email)
	if err != nil {
		return
	}
	if user == nil {
		return nil, model.ErrUserNotFound
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))
	if err != nil {
		return nil, errs.Wrap("compare password and hash", err)
	}

	userDTO := model.ParseUserDTO(*user)
	tokens, err := us.manager.Service().Auth().NewPairTokens(userDTO)
	if err != nil {
		return
	}

	result = &model.AuthResponse{
		User:   userDTO,
		Tokens: *tokens,
	}
	return
}

func (us *UserService) Refresh(ctx context.Context, refreshToken string) (result *model.Tokens, err error) {
	prevParams, err := us.manager.Repository().UserParams().FindByToken(ctx, refreshToken)
	if err != nil {
		return
	}

	if prevParams == nil {
		return nil, model.ErrTokenNotFound
	}

	token, err := jwt.ParseWithClaims(prevParams.RefreshToken, &jwt.RegisteredClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(viper.GetString("JWT_REFRESH_SECRET_KEY")), nil
	})
	if err != nil {
		return nil, errs.Wrap("parse token", err)
	}

	user := &model.UserDTO{}
	claims, ok := token.Claims.(*jwt.RegisteredClaims)

	err = json.Unmarshal([]byte(claims.Issuer), user)
	if err != nil {
		return nil, errs.Wrap("unmarshal issuer", err)
	}
	if !ok || !token.Valid || user == nil {
		return nil, model.ErrInvalidToken
	}

	result, err = us.manager.Service().Auth().NewPairTokens(*user)
	if err != nil {
		return
	}

	if err := us.manager.Service().User().SaveToken(ctx, (*prevParams).TokenID, (*result).RefreshToken); err != nil {
		return nil, errs.Wrap("service.user.SaveToken", err)
	}

	return
}

func (us *UserService) Logout(ctx context.Context, refreshToken string) (err error) {
	result, err := us.manager.Repository().UserParams().DeleteToken(ctx, refreshToken)
	if err != nil {
		return
	}
	if result == nil {
		return model.ErrTokenNotFound
	}
	return
}

func (us *UserService) SaveToken(ctx context.Context, ID int64, refreshToken string) (err error) {
	prevParams, err := us.manager.Repository().UserParams().FindByID(ctx, ID)
	if err != nil {
		return
	}

	if prevParams != nil {
		return us.manager.Repository().UserParams().Update(ctx, ID, refreshToken)
	}

	_, err = us.manager.Repository().UserParams().Create(ctx, ID, refreshToken)
	return
}
