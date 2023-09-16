package postgres

import (
	"app/internal/config"
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

func getDSN(cfg *config.Cfg) string {
	return fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		cfg.PostgresHost, cfg.PostgresPort, cfg.PostgresUser, cfg.PostgresPassword, cfg.PostgresDatabase, cfg.PostgresSSL)
}

func InitConnection(ctx context.Context, cfg *config.Cfg) *pgxpool.Pool {
	connConfig, err := pgxpool.ParseConfig(getDSN(cfg))
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
