package interfaces

import (
	"app/internal/model"
	"context"
)

type IService interface {
	Auth() IAuthService
	User() IUserService
	Storage() IStorageService
}

type IAuthService interface {
	NewPairTokens(user model.UserDTO) (result *model.Tokens, err error)
}

type IUserService interface {
	Create(ctx context.Context, body model.SignUp) (result *model.AuthResponse, err error)
	SaveToken(ctx context.Context, ID int64, refreshToken string) (err error)
	Authenticate(ctx context.Context, body model.SignIn) (result *model.AuthResponse, err error)
	Refresh(ctx context.Context, refreshToken string) (result *model.Tokens, err error)
	Logout(ctx context.Context, refreshToken string) (err error)
}

type IStorageService interface {
	SaveFile(ctx context.Context, obj *model.HandlePicture) (string, error)
}
