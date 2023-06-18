package service

import (
	"app/internal/model"
	"context"
)

type IStorageService interface {
	HandlePicture(ctx context.Context, obj *model.HandlePictureReq) (string, error)
}
