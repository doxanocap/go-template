package config

import (
	"github.com/doxanocap/pkg/env"
	"github.com/doxanocap/pkg/lg"
	"github.com/spf13/viper"
)

type Cfg struct {
	AppVersion     string `mapstructure:"APP_VERSION"`
	AppEnvironment string `mapstructure:"APP_ENVIRONMENT"`

	ServerIP   string `mapstructure:"SERVER_IP"`
	ServerPort string `mapstructure:"SERVER_PORT"`

	RefreshTokenKey string `mapstructure:"REFRESH_TOKEN_KEY"`
	AccessTokenKey  string `mapstructure:"ACCESS_TOKEN_KEY"`

	LogJSON bool `mapstructure:"LOG_JSON"`

	PostgresHost     string `mapstructure:"PSQL_HOST"`
	PostgresPort     string `mapstructure:"PSQL_PORT"`
	PostgresUser     string `mapstructure:"PSQL_USER"`
	PostgresPassword string `mapstructure:"PSQL_PASSWORD"`
	PostgresDatabase string `mapstructure:"PSQL_DATABASE"`
	PostgresSSL      string `mapstructure:"PSQL_SSL"`

	RabbitDSN string `mapstructure:"RABBIT_DSN"`

	RedisPrefix   string `mapstructure:"REDIS_PREFIX"`
	RedisHost     string `mapstructure:"REDIS_HOST"`
	RedisPassword string `mapstructure:"REDIS_PASSWORD"`
	RedisDatabase int    `mapstructure:"REDIS_DATABASE"`

	AwsS3AccessKey   string `mapstructure:"S3_ACCESS_KEY"`
	AwsS3SecretKey   string `mapstructure:"S3_SECRET_KEY"`
	AwsS3EndpointUrl string `mapstructure:"S3_ENDPOINT_URL"`
	AwsS3BucketName  string `mapstructure:"S3_BUCKET_NAME"`
	AwsS3UseSSL      bool   `mapstructure:"S3_USE_SSL"`
}

func InitConfig() *Cfg {
	config := &Cfg{}
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

	return config
}
