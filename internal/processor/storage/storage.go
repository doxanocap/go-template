package storage

import (
	"app/internal/manager/interfaces/processor"
	"context"
	"fmt"
	"io"
)

type StorageProcessor struct {
	provider processor.IStorageProvider
}

func (sp *StorageProcessor) Save(ctx context.Context, name string, file io.Reader, size int64) error {
	fmt.Println(sp.provider)
	return sp.provider.Save(ctx, name, file, size)
}

func Init(provider processor.IStorageProvider) *StorageProcessor {
	return &StorageProcessor{
		provider: provider,
	}
}
