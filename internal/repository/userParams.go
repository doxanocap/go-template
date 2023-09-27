package repository

import (
	"app/internal/model"
	"context"

	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/doxanocap/pkg/errs"
	"github.com/georgysavva/scany/pgxscan"
	"github.com/jackc/pgx/v4/pgxpool"
	"go.uber.org/zap"
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
	query := repo.builder.
		Insert("user_params").
		Columns(
			"user_params",
			"refresh_token",
			"updated_at").
		Values(
			ID,
			refreshToken,
			time.Now(),
		).
		Suffix("RETURNING *")

	raw, args := query.MustSql()
	err = pgxscan.Select(ctx, repo.pool, result, raw, args...)
	if err != nil {
		return nil, errs.Wrap("repository.user_params.Create", err)
	}

	return
}

func (repo *UserParamsRepository) FindByID(ctx context.Context, ID int64) (result *model.UserParams, err error) {
	query := repo.builder.
		Select("*").
		From("user_params").
		Where(sq.Eq{"token_id": ID})

	raw, args := query.MustSql()
	err = pgxscan.Select(ctx, repo.pool, result, raw, args...)
	if err != nil {
		return nil, errs.Wrap("repository.user_params.FindByID", err)
	}

	return
}

func (repo *UserParamsRepository) FindByToken(ctx context.Context, refreshToken string) (result *model.UserParams, err error) {
	result = &model.UserParams{}
	query := repo.builder.
		Select("*").
		From("user_params").
		Where(sq.Eq{"refresh_token": refreshToken})

	raw, args := query.MustSql()
	err = pgxscan.Select(ctx, repo.pool, result, raw, args...)
	if err != nil {
		return nil, errs.Wrap("repository.user_params.FindByToken", err)
	}

	return
}

func (repo *UserParamsRepository) Update(ctx context.Context, ID int64, refreshToken string) (err error) {
	query := repo.builder.
		Update("user_params").
		SetMap(map[string]interface{}{
			"refresh_token": refreshToken,
			"updated_at":    time.Now(),
		}).
		Where(sq.Eq{"token_id": ID})

	raw, args := query.MustSql()
	_, err = repo.pool.Exec(ctx, raw, args...)
	if err != nil {
		return errs.Wrap("repository.user_params.Update", err)
	}

	return
}

func (repo *UserParamsRepository) DeleteToken(ctx context.Context, refreshToken string) (result *model.UserParams, err error) {
	result = &model.UserParams{}
	query := repo.builder.
		Delete("user_params").
		Where(sq.Eq{"refresh_token": refreshToken}).
		Suffix("RETURNING *")

	raw, args := query.MustSql()
	err = pgxscan.Select(ctx, repo.pool, result, raw, args...)
	if err != nil {
		return nil, errs.Wrap("repository.user_params.DeleteToken", err)
	}

	return
}
