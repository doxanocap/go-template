package repository

import (
	"app/internal/model"
	"context"
	sq "github.com/Masterminds/squirrel"
	"github.com/georgysavva/scany/pgxscan"
	"github.com/jackc/pgx/v4/pgxpool"
	"go.uber.org/zap"
)

const (
	userTable         = "user"
	userRowID         = "id"
	emailColumn       = "email"
	usernameColumn    = "username"
	phoneNumberColumn = "phone_number"
	passwordColumn    = "password"
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

func (repo *UserRepository) Create(ctx context.Context, obj model.SignUp) (*model.User, error) {
	result := &model.User{}
	log := repo.log.Named("Create").With(
		zap.String(emailColumn, obj.Email),
		zap.String(usernameColumn, obj.Username),
		zap.String(phoneNumberColumn, obj.PhoneNumber))

	query := repo.builder.
		Insert(userTable).
		Columns(
			emailColumn,
			usernameColumn,
			phoneNumberColumn,
			passwordColumn).
		Values(
			obj.Email,
			obj.Username,
			obj.PhoneNumber,
			obj.Password).
		Suffix("RETURNING *")

	raw, args := query.MustSql()
	log.Info("query", zap.String("raw", raw), zap.Any("args", args))

	err := pgxscan.Select(ctx, repo.pool, result, raw, args...)
	if err != nil {
		log.Error("failed", zap.Error(err))
	}

	return result, err
}

func (repo *UserRepository) FindByEmail(ctx context.Context, email string) (*model.User, error) {
	result := &model.User{}
	log := repo.log.Named("FindByEmail").With(
		zap.Any(emailColumn, email))

	query := repo.builder.
		Select("*").
		From(userTable).
		Where(sq.Eq{emailColumn: email})

	raw, args := query.MustSql()
	log.Info("query", zap.String("raw", raw), zap.Any("args", args))

	err := pgxscan.Select(ctx, repo.pool, result, raw, args...)
	if err != nil {
		log.Error("failed", zap.Error(err))
	}

	return result, err
}
