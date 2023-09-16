package config

import (
	"github.com/doxanocap/pkg/env"
	"github.com/doxanocap/pkg/lg"
	"github.com/spf13/viper"
)

type Cfg struct {
	AppVersion     string `env:"APP_VERSION"`
	AppEnvironment string `env:"APP_ENVIRONMENT"`

	ServerIP   string `env:"SERVER_IP"`
	ServerPort string `env:"SERVER_PORT"`

	RefreshTokenKey string `env:"REFRESH_TOKEN_KEY"`
	AccessTokenKey  string `env:"ACCESS_TOKEN_KEY"`

	LogJSON bool `env:"LOG_JSON"`

	PostgresHost     string `env:"PSQL_HOST"`
	PostgresPort     string `env:"PSQL_PORT"`
	PostgresUser     string `env:"PSQL_USER"`
	PostgresPassword string `env:"PSQL_PASSWORD"`
	PostgresDatabase string `env:"PSQL_DATABASE"`
	PostgresSSL      string `env:"PSQL_SSL"`

	RabbitDSN string `env:"RABBIT_DSN"`

	RedisPrefix   string `env:"REDIS_PREFIX"`
	RedisHost     string `env:"REDIS_HOST"`
	RedisPassword string `env:"REDIS_PASSWORD"`
	RedisDatabase int    `env:"REDIS_DATABASE"`
}

func InitConfig() *Cfg {
	config := &Cfg{}
	env.SetDefaultTag("env")
	env.SetViperDefaults(config)

	viper.SetConfigType("env")
	viper.SetConfigFile(".env")
	if err := viper.ReadInConfig(); err != nil {
		lg.Infof("config: %s", err)
	}

	viper.AutomaticEnv()
	err := viper.Unmarshal(&config)
	if err != nil {
		lg.Fatalf("unmarshal config: %s", err)
	}

	viper.AutomaticEnv()
	return config
}
