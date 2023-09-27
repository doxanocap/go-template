package repository

import (
	"app/internal/model"
	"context"

	sq "github.com/Masterminds/squirrel"
	"github.com/doxanocap/pkg/errs"
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
	query := repo.builder.
		Insert("user").
		Columns(
			"email",
			"username",
			"phone_number",
			"password").
		Values(
			obj.Email,
			obj.Username,
			obj.PhoneNumber,
			obj.Password).
		Suffix("RETURNING *")

	raw, args := query.MustSql()

	err = pgxscan.Select(ctx, repo.pool, result, raw, args...)
	if err != nil {
		return nil, errs.Wrap("repository.user.Create", err)
	}

	return
}

func (repo *UserRepository) FindByUUID(ctx context.Context, uuid string) (result *model.User, err error) {
	result = &model.User{}
	query := repo.builder.
		Select("*").
		From("user").
		Where(sq.Eq{"uuid": uuid})

	raw, args := query.MustSql()

	err = pgxscan.Select(ctx, repo.pool, result, raw, args...)
	if err != nil {
		return nil, errs.Wrap("repository.user.FindByUUID", err)
	}

	return
}

func (repo *UserRepository) FindByEmail(ctx context.Context, email string) (result *model.User, err error) {
	result = &model.User{}
	query := repo.builder.
		Select("*").
		From("user").
		Where(sq.Eq{"email": email})

	raw, args := query.MustSql()
	err = pgxscan.Select(ctx, repo.pool, result, raw, args...)
	if err != nil {
		return nil, errs.Wrap("repository.user.FindByEmail", err)
	}

	return
}
