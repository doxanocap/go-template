package postgres

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v4/pgxpool"
	_ "github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

const (
	driver        = "postgres"
	ChatListTable = "chat_list"
	ChatMessages  = "chat_messages"
)

type Config struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
	SSLMode  string
}

func getDSN(cfg Config) string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s", cfg.Host, cfg.Port, cfg.Username, cfg.Password, cfg.DBName, cfg.SSLMode)
}

func Connect(ctx context.Context, dsn string) (*pgxpool.Pool, error) {
	connConfig, err := pgxpool.ParseConfig(dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to parse config -> %v", err)
	}

	conn, err := pgxpool.ConnectConfig(ctx, connConfig)
	if err != nil {
		return nil, fmt.Errorf("failed to connect -> %v", err)
	}

	err = conn.Ping(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to ping -> %v", err)
	}

	return conn, nil
}
