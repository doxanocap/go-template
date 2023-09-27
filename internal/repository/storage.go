package repository

import (
	"app/internal/model"
	"context"

	"github.com/Masterminds/squirrel"
	"github.com/doxanocap/pkg/errs"
	"github.com/georgysavva/scany/pgxscan"
	"github.com/jackc/pgx/v4/pgxpool"
	"go.uber.org/zap"
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
	query := repo.builder.
		Insert("storage").
		Columns("key", "format").
		Values(key, format).
		Suffix("RETURNING *")

	raw, args := query.MustSql()
	err = pgxscan.Get(ctx, repo.pool, result, raw, args...)
	if err != nil {
		return nil, errs.Wrap("repository.storage.Create", err)
	}

	return
}
