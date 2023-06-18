package service

import (
	"app/internal/cns"
	"app/internal/cns/errs"
	"app/internal/manager/interfaces"
	"app/internal/model"
	"context"
	"fmt"
	"strings"
)

const (
	dot = "."
)

type StorageService struct {
	repository interfaces.IRepository
	processor  interfaces.IProcessor
}

func InitStorageService(repository interfaces.IRepository, processor interfaces.IProcessor) *StorageService {
	return &StorageService{
		repository: repository,
		processor:  processor,
	}
}

func (ss *StorageService) HandlePicture(ctx context.Context, obj *model.HandlePictureReq) (string, error) {
	pictureFormat := ss.getPictureFormat(obj.FileName)
	if !cns.IsValidFormat(pictureFormat) {
		return cns.NilString, errs.InvalidFormat()
	}

	storage, err := ss.repository.Storage().Create(ctx, obj.Key, pictureFormat)
	if err != nil {
		return cns.NilString, err
	}

	obj.FileName = fmt.Sprintf("%s_%d.%s", obj.Key, (*storage).ID, pictureFormat)

	err = ss.processor.Storage().Save(ctx, obj.FileName, obj.File, obj.Size)
	if err != nil {
		return cns.NilString, err
	}

	return obj.FileName, nil
}

func (ss *StorageService) getPictureFormat(filename string) string {
	divided := strings.Split(filename, dot)
	if len(divided) != 2 {
		return cns.NilString
	}
	return divided[1]
}
