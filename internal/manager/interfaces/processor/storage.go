package processor

import (
	"context"
	"io"
)

type IStorageProcessor interface {
	Save(ctx context.Context, name string, file io.Reader, size int64) error
}

type IStorageProvider interface {
	Save(ctx context.Context, name string, file io.Reader, size int64) error
}
