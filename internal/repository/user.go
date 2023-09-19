package repository

import (
	"app/internal/cns"
	"app/internal/model"
	"context"
	sq "github.com/Masterminds/squirrel"
	"github.com/georgysavva/scany/pgxscan"
	"github.com/jackc/pgx/v4/pgxpool"
	"go.uber.org/zap"
)

type UserRepository struct {
	pool    *pgxpool.Pool
	log     *zap.Logger
	builder sq.StatementBuilderType
}

func InitUserRepository(pool *pgxpool.Pool, log *zap.Logger) *UserRepository {
	return &UserRepository{
		pool:    pool,
		log:     log,
		builder: sq.StatementBuilder.PlaceholderFormat(sq.Dollar),
	}
}

func (repo *UserRepository) Create(ctx context.Context, obj model.SignUp) (result *model.User, err error) {
	result = &model.User{}
	log := repo.log.Named("Create").With(
		zap.String(cns.EmailColumn, obj.Email),
		zap.String(cns.UsernameColumn, obj.Username),
		zap.String(cns.PhoneNumberColumn, obj.PhoneNumber))

	query := repo.builder.
		Insert(cns.UserTable).
		Columns(
			cns.EmailColumn,
			cns.UsernameColumn,
			cns.PhoneNumberColumn,
			cns.PasswordColumn).
		Values(
			obj.Email,
			obj.Username,
			obj.PhoneNumber,
			obj.Password).
		Suffix("RETURNING *")

	raw, args := query.MustSql()
	log.Info("query", zap.String("raw", raw), zap.Any("args", args))

	err = pgxscan.Select(ctx, repo.pool, result, raw, args...)
	if err != nil {
		log.Error("failed", zap.Error(err))
	}

	return
}

func (repo *UserRepository) FindByUUID(ctx context.Context, uuid string) (result *model.User, err error) {
	result = &model.User{}
	log := repo.log.Named("FindByID").With(
		zap.String(cns.UserTableUUID, uuid))

	query := repo.builder.
		Select("*").
		From(cns.UserTable).
		Where(sq.Eq{cns.UserTableUUID: uuid})

	raw, args := query.MustSql()
	log.Info("query", zap.String("raw", raw), zap.Any("args", args))

	err = pgxscan.Select(ctx, repo.pool, result, raw, args...)
	if err != nil {
		log.Error("failed", zap.Error(err))
	}

	return
}

func (repo *UserRepository) FindByEmail(ctx context.Context, email string) (result *model.User, err error) {
	result = &model.User{}
	log := repo.log.Named("FindByEmail").With(
		zap.String(cns.EmailColumn, email))

	query := repo.builder.
		Select("*").
		From(cns.UserTable).
		Where(sq.Eq{cns.EmailColumn: email})

	raw, args := query.MustSql()
	log.Info("query", zap.String("raw", raw), zap.Any("args", args))

	err = pgxscan.Select(ctx, repo.pool, result, raw, args...)
	if err != nil {
		log.Error("failed", zap.Error(err))
	}

	return
}
