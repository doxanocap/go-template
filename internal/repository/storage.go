package repository

import (
	"app/internal/cns/errs"
	"app/internal/model"
	"context"
	"github.com/Masterminds/squirrel"
	"github.com/georgysavva/scany/pgxscan"
	"github.com/jackc/pgx/v4/pgxpool"
	"go.uber.org/zap"
)

const (
	storageTable   = "storage"
	storageTableID = "id"
	keyColumn      = "key"
	formatColumn   = "format"
)

type StorageRepository struct {
	pool    *pgxpool.Pool
	log     *zap.Logger
	builder squirrel.StatementBuilderType
}

func InitStorageRepository(pool *pgxpool.Pool, log *zap.Logger) *StorageRepository {
	return &StorageRepository{
		pool:    pool,
		log:     log,
		builder: squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar),
	}
}

func (repo *StorageRepository) Create(ctx context.Context, key, format string) (result *model.Storage, err error) {
	result = &model.Storage{}
	log := repo.log.Named("Create").With(
		zap.String(keyColumn, key),
		zap.String(formatColumn, format))

	query := repo.builder.
		Insert(storageTable).
		Columns(keyColumn, formatColumn).
		Values(key, format).
		Suffix("RETURNING *")

	raw, args := query.MustSql()
	log.Info("query", zap.String("raw", raw), zap.Any("args", args))

	err = pgxscan.Get(ctx, repo.pool, result, raw, args...)
	if err != nil {
		log.Error("failed", zap.Error(err))
		return
	}
	if result == nil {
		return nil, errs.EmptyResult()
	}
	return
}
