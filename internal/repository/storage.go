package repository

import (
	"app/internal/cns/errs"
	"app/internal/model"
	"context"
	"github.com/Masterminds/squirrel"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

const (
	storageTable   = "storage"
	storageTableID = "id"
	keyColumn      = "key"
	formatColumn   = "format"
)

type StorageRepository struct {
	conn    *gorm.DB
	log     *zap.Logger
	builder squirrel.StatementBuilderType
}

func InitStorageRepository(conn *gorm.DB, log *zap.Logger) *StorageRepository {
	return &StorageRepository{
		conn:    conn,
		log:     log,
		builder: squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar),
	}
}

func (repo *StorageRepository) Create(ctx context.Context, key, format string) (result *model.Storage, err error) {
	result = &model.Storage{}
	log := repo.log.Named("Create").With(
		zap.String(keyColumn, key),
		zap.String(formatColumn, format))

	repo.conn.Create(result)
	log.Info("query")

	if result == nil {
		return nil, errs.EmptyResult()
	}
	return
}
