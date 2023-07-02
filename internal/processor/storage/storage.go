package storage

import (
	"app/internal/cns"
	"app/internal/manager/interfaces/processor"
	"context"
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
	s.log.Named("Save").With(
		zap.String(cns.StorageFileName, name),
		zap.Int64(cns.StorageFileSize, size)).Info("saved")

	return s.provider.Save(ctx, name, file, size)
}
