package repository

import (
	"app/internal/cns"
	"app/internal/model"
	"context"
	sq "github.com/Masterminds/squirrel"
	"github.com/georgysavva/scany/pgxscan"
	"github.com/jackc/pgx/v4/pgxpool"
	"go.uber.org/zap"
	"time"
)

type UserParamsRepository struct {
	pool    *pgxpool.Pool
	log     *zap.Logger
	builder sq.StatementBuilderType
}

func InitUserParamsRepository(pool *pgxpool.Pool, log *zap.Logger) *UserParamsRepository {
	return &UserParamsRepository{
		pool:    pool,
		log:     log,
		builder: sq.StatementBuilder.PlaceholderFormat(sq.Dollar),
	}
}

func (repo *UserParamsRepository) Create(ctx context.Context, ID int64, refreshToken string) (result *model.UserParams, err error) {
	result = &model.UserParams{}
	log := repo.log.Named("Create").With(
		zap.Int64(cns.UserParamsTableID, ID),
		zap.String(cns.RefreshTokenColumn, refreshToken))

	query := repo.builder.
		Insert(cns.UserParamsTable).
		Columns(
			cns.UserParamsTableID,
			cns.RefreshTokenColumn,
			cns.UpdatedAtColumn).
		Values(
			ID,
			refreshToken,
			time.Now(),
		).
		Suffix("RETURNING *")

	raw, args := query.MustSql()
	log.Info("query", zap.String("raw", raw), zap.Any("args", args))

	err = pgxscan.Select(ctx, repo.pool, result, raw, args...)
	if err != nil {
		log.Error("failed", zap.Error(err))
	}

	return
}

func (repo *UserParamsRepository) FindByID(ctx context.Context, ID int64) (result *model.UserParams, err error) {
	result = &model.UserParams{}
	log := repo.log.Named("FindByID").With(
		zap.Int64(cns.UserParamsTableID, ID))

	query := repo.builder.
		Select("*").
		From(cns.UserParamsTable).
		Where(sq.Eq{cns.UserParamsTableID: ID})

	raw, args := query.MustSql()
	log.Info("query", zap.String("raw", raw), zap.Any("args", args))

	err = pgxscan.Select(ctx, repo.pool, result, raw, args...)
	if err != nil {
		log.Error("failed", zap.Error(err))
	}

	return
}

func (repo *UserParamsRepository) FindByToken(ctx context.Context, refreshToken string) (result *model.UserParams, err error) {
	result = &model.UserParams{}
	log := repo.log.Named("FindByToken").With(
		zap.String(cns.RefreshTokenColumn, refreshToken))

	query := repo.builder.
		Select("*").
		From(cns.UserParamsTable).
		Where(sq.Eq{cns.RefreshTokenColumn: refreshToken})

	raw, args := query.MustSql()
	log.Info("query", zap.String("raw", raw), zap.Any("args", args))

	err = pgxscan.Select(ctx, repo.pool, result, raw, args...)
	if err != nil {
		log.Error("failed", zap.Error(err))
	}

	return
}

func (repo *UserParamsRepository) Update(ctx context.Context, ID int64, refreshToken string) (err error) {
	log := repo.log.Named("FindByToken").With(
		zap.Int64(cns.UserParamsTableID, ID),
		zap.String(cns.RefreshTokenColumn, refreshToken))

	query := repo.builder.
		Update(cns.UserParamsTable).
		SetMap(map[string]interface{}{
			cns.RefreshTokenColumn: refreshToken,
			cns.UpdatedAtColumn:    time.Now(),
		}).
		Where(sq.Eq{cns.UserParamsTableID: ID})

	raw, args := query.MustSql()
	log.Info("query", zap.String("raw", raw), zap.Any("args", args))

	_, err = repo.pool.Exec(ctx, raw, args...)
	if err != nil {
		log.Error("failed", zap.Error(err))
	}

	return
}

func (repo *UserParamsRepository) DeleteToken(ctx context.Context, refreshToken string) (result *model.UserParams, err error) {
	result = &model.UserParams{}
	log := repo.log.Named("FindByToken").With(
		zap.String(cns.RefreshTokenColumn, refreshToken))

	query := repo.builder.
		Delete(cns.UserParamsTable).
		Where(sq.Eq{cns.RefreshTokenColumn: refreshToken}).
		Suffix("RETURNING *")

	raw, args := query.MustSql()
	log.Info("query", zap.String("raw", raw), zap.Any("args", args))

	_, err = repo.pool.Exec(ctx, raw, args...)
	if err != nil {
		log.Error("failed", zap.Error(err))
	}

	err = pgxscan.Select(ctx, repo.pool, result, raw, args...)
	if err != nil {
		log.Error("failed", zap.Error(err))
	}

	return
}
