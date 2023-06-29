package storage

import (
	"app/internal/manager/interfaces/processor"
	"context"
	"io"
)

type Storage struct {
	provider processor.IStorageProvider
}

func (s *Storage) Save(ctx context.Context, name string, file io.Reader, size int64) error {
	return s.provider.Save(ctx, name, file, size)
}

func Init(provider processor.IStorageProvider) *Storage {
	return &Storage{
		provider: provider,
	}
}
