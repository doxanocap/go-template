package postgres

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

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

func Connect(dsn string) (*gorm.DB, error) {
	connection, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("failed to connect to postgres -> %v", err)
	}

	sqlDB, err := connection.DB()
	if err != nil {
		return nil, fmt.Errorf("failed to connect to postgres -> %v", err)
	}

	err = sqlDB.Ping()
	if err != nil {
		return nil, fmt.Errorf("failed to ping -> %v", err)
	}
	sqlDB.SetMaxIdleConns(100)
	sqlDB.SetMaxOpenConns(100)

	return connection, nil
}
