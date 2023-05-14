package repository

import (
	"github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v4/pgxpool"
	"go.uber.org/zap"
)

type AuthRepository struct {
	pool    *pgxpool.Pool
	log     *zap.Logger
	builder squirrel.StatementBuilderType
}

func InitAuthRepository(pool *pgxpool.Pool, log *zap.Logger) *AuthRepository {
	return &AuthRepository{
		pool:    pool,
		log:     log,
		builder: squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar),
	}
}
