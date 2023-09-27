package service

import (
	"app/internal/consts"
	"app/internal/manager/interfaces"
	"app/internal/model"
	"context"
	"fmt"
	"github.com/doxanocap/pkg/errs"
	"strings"
)

type StorageService struct {
	manager interfaces.IManager
}

func InitStorageService(manager interfaces.IManager) *StorageService {
	return &StorageService{
		manager: manager,
	}
}

func (ss *StorageService) SaveFile(ctx context.Context, obj *model.HandlePicture) (string, error) {
	pictureFormat := ss.getPictureFormat(obj.FileName)
	if !consts.IsValidFormat(pictureFormat) {
		return consts.NilString, model.ErrInvalidFileFormat
	}

	storage, err := ss.manager.Repository().Storage().Create(ctx, obj.Key, pictureFormat)
	if err != nil {
		return consts.NilString, err
	}

	obj.FileName = fmt.Sprintf("%s_%d.%s", obj.Key, (*storage).ID, pictureFormat)

	err = ss.manager.Processor().Storage().Save(ctx, obj.FileName, obj.File, obj.Size)
	if err != nil {
		return consts.NilString, errs.Wrap("processor.storage.Save", err)
	}

	return obj.FileName, nil
}

func (ss *StorageService) getPictureFormat(filename string) string {
	divided := strings.Split(filename, ".")
	if len(divided) != 2 {
		return consts.NilString
	}
	return divided[1]
}
