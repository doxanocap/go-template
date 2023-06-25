package storage

import (
	"app/internal/manager/interfaces/processor"
	"context"
	"io"
)

type StorageProcessor struct {
	provider processor.IStorageProvider
}

func (sp *StorageProcessor) Save(ctx context.Context, name string, file io.Reader, size int64) error {
	return sp.provider.Save(ctx, name, file, size)
}

func Init(provider processor.IStorageProvider) *StorageProcessor {
	return &StorageProcessor{
		provider: provider,
	}
}
