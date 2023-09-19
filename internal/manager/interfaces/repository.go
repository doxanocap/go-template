package interfaces

import (
	"app/internal/model"
	"context"
)

type IRepository interface {
	User() IUserRepository
	UserParams() IUserParamsRepository
	Storage() IStorageRepository
}

type IStorageRepository interface {
	Create(ctx context.Context, key, format string) (res *model.Storage, err error)
}

type IUserRepository interface {
	Create(ctx context.Context, obj model.SignUp) (result *model.User, err error)
	FindByUUID(ctx context.Context, uuid string) (result *model.User, err error)
	FindByEmail(ctx context.Context, email string) (result *model.User, err error)
}

type IUserParamsRepository interface {
	Create(ctx context.Context, ID int64, refreshToken string) (result *model.UserParams, err error)
	FindByID(ctx context.Context, ID int64) (result *model.UserParams, err error)
	FindByToken(ctx context.Context, refreshToken string) (result *model.UserParams, err error)
	Update(ctx context.Context, ID int64, refreshToken string) (err error)
	DeleteToken(ctx context.Context, refreshToken string) (result *model.UserParams, err error)
}
