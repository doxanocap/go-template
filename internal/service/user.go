package service

import (
	"app/internal/cns/errs"
	"app/internal/manager/interfaces"
	"app/internal/model"
	"context"
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
