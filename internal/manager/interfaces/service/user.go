package service

import (
	"app/internal/model"
	"context"
)

type IUserService interface {
	Create(ctx context.Context, body model.SignUp) (result *model.AuthResponse, err error)
	Authenticate(ctx context.Context, body model.SignIn) (result *model.AuthResponse, err error)
}
