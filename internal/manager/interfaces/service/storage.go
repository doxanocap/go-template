package service

import (
	"app/internal/model"
	"context"
)

type IStorageService interface {
	SaveFile(ctx context.Context, obj *model.HandlePicture) (string, error)
}
