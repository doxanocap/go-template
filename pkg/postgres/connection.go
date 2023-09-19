package postgres

import (
	"app/internal/model"
	"context"
	"fmt"
	"github.com/doxanocap/pkg/lg"

	"github.com/jackc/pgx/v4/pgxpool"
	_ "github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

const (
	driver        = "postgres"
	ChatListTable = "chat_list"
	ChatMessages  = "chat_messages"
)

func getDSN(cfg model.Psql) string {
	return fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		cfg.PsqlHost, cfg.PsqlPort, cfg.PsqlUser, cfg.PsqlPassword, cfg.PsqlDatabase, cfg.PsqlSSL)
}

func InitConnection(cfg *model.Config) *pgxpool.Pool {
	ctx := context.Background()
	connConfig, err := pgxpool.ParseConfig(getDSN(cfg.Psql))
	if err != nil {
		lg.Fatalf("failed to parse config -> %v", err)
	}

	conn, err := pgxpool.ConnectConfig(ctx, connConfig)
	if err != nil {
		lg.Fatalf("failed to connect -> %v", err)
	}

	err = conn.Ping(ctx)
	if err != nil {
		lg.Fatalf("failed to ping -> %v", err)
	}

	return conn
}
