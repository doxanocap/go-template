package service

import (
	"app/internal/cns/errs"
	"app/internal/manager/interfaces"
	"app/internal/model"
	"context"
	"encoding/json"
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
		err = errs.HttpConflict("already exists")
		return
	}
	salt, _ := bcrypt.GenerateFromPassword([]byte(""), 10)

	passwordWithSalt := append(body.Password, salt...)
	hashedPassword, err := bcrypt.GenerateFromPassword(passwordWithSalt, bcrypt.DefaultCost)
	if err != nil {
		return
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
		err = errs.HttpNotFound("user")
		return
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))
	if err != nil {
		return
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
		err = errs.HttpNotFound("token")
		return
	}

	token, err := jwt.ParseWithClaims(prevParams.RefreshToken, &jwt.RegisteredClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(viper.GetString("JWT_REFRESH_SECRET_KEY")), nil
	})
	if err != nil {
		return
	}

	user := &model.UserDTO{}
	claims, ok := token.Claims.(*jwt.RegisteredClaims)

	err = json.Unmarshal([]byte(claims.Issuer), user)
	if err != nil {
		return
	}
	if !ok || !token.Valid || user == nil {
		err = errs.HttpConflict("invalid token")
		return
	}

	result, err = us.manager.Service().Auth().NewPairTokens(*user)
	if err != nil {
		return
	}

	if err := us.manager.Service().User().SaveToken(ctx, (*prevParams).TokenID, (*result).RefreshToken); err != nil {
		return nil, err
	}

	return
}

func (us *UserService) Logout(ctx context.Context, refreshToken string) (err error) {
	result, err := us.manager.Repository().UserParams().DeleteToken(ctx, refreshToken)
	if err != nil {
		return
	}
	if result == nil {
		return errs.HttpNotFound("token")
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
