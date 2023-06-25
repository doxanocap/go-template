package repository

import (
	"app/internal/model"
	"context"
)

type IStorageRepository interface {
	Create(ctx context.Context, key, format string) (res *model.Storage, err error)
}
