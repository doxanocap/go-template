package storage

import (
	"app/internal/manager/interfaces/processor"
	"context"
	"fmt"
	"github.com/doxanocap/pkg/errs"
	"go.uber.org/zap"
	"io"
)

type Storage struct {
	provider processor.IStorageProvider
	log      *zap.Logger
}

func Init(provider processor.IStorageProvider, log *zap.Logger) *Storage {
	return &Storage{
		provider: provider,
		log:      log,
	}
}

func (s *Storage) Save(ctx context.Context, name string, file io.Reader, size int64) error {
	log := s.log.Named("[Save]")

	err := s.provider.Save(ctx, name, file, size)
	if err != nil {
		log.Error(fmt.Sprintf("storage provider: %s", err))
		err = errs.Wrap("processor.storage.Save", err)
	}
	return err
}
