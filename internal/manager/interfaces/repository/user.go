package repository

import (
	"app/internal/model"
	"context"
)

type IUserRepository interface {
	Create(ctx context.Context, obj model.SignUp) (result *model.User, err error)
	FindByEmail(ctx context.Context, email string) (result *model.User, err error)
}
