package repository

import (
	"app/internal/model"
	"context"
)

type IUserParamsRepository interface {
	Create(ctx context.Context, ID int64, refreshToken string) (result *model.UserParams, err error)
	FindByID(ctx context.Context, ID int64) (result *model.UserParams, err error)
	FindByToken(ctx context.Context, refreshToken string) (result *model.UserParams, err error)
	Update(ctx context.Context, ID int64, refreshToken string) (err error)
	DeleteToken(ctx context.Context, refreshToken string) (result *model.UserParams, err error)
}
