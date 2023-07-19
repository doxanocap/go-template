package cmd

import (
	"fmt"
	"github.com/spf13/viper"
)

var Conf = struct {
	Debug          bool   `mapstructure:"DEBUG"`
	LogLevel       string `mapstructure:"LOG_LEVEL"`
	HttpListen     string `mapstructure:"HTTP_LISTEN"`
	HttpCors       bool   `mapstructure:"HTTP_CORS"`
	SwagHost       string `mapstructure:"SWAG_HOST"`
	SwagBasePath   string `mapstructure:"SWAG_BASE_PATH"`
	SwagSchema     string `mapstructure:"SWAG_SCHEMA"`
	PgDsn          string `mapstructure:"PG_DSN"`
	RedisUrl       string `mapstructure:"REDIS_URL"`
	RedisPsw       string `mapstructure:"REDIS_PSW"`
	RedisDb        int    `mapstructure:"REDIS_DB"`
	RedisKeyPrefix string `mapstructure:"REDIS_KEY_PREFIX"`
	MsJwtsUrl      string `mapstructure:"MS_JWTS_URL"`
	MsWsUrl        string `mapstructure:"MS_WS_URL"`
	MsPushUrl      string `mapstructure:"MS_PUSH_URL"`
	NoSmsCheck     bool   `mapstructure:"NO_SMS_CHECK"`
	SMSCLogin      string `mapstructure:"SMSC_LOGIN"`
	SMSCPassword   string `mapstructure:"SMSC_PASSWORD"`
	MobizonApiKey  string `mapstructure:"MOBIZON_API_KEY"`
	JWTSecret      string `mapstructure:"JWT_SECRET"`
}{}

func InitConfig() {
	viper.SetDefault("ENV_MODE", "development")

	viper.SetDefault("ZAP_JSON", false)

	viper.SetDefault("TOKEN_MAX_AGE", 30*24*60*60*1000)
	viper.SetDefault("TOKEN_PATH", "/")
	viper.SetDefault("TOKEN_DOMAIN", "localhost")

	viper.SetDefault("JWT_REFRESH_SECRET_KEY", "secret1")
	viper.SetDefault("JWT_ACCESS_SECRET_KEY", "secret1")

	viper.SetDefault("RABBIT_MQ_DSN", "amqp://guest:guest@localhost:5672")

	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AddConfigPath("./config")

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}

	viper.AutomaticEnv()

	err = viper.Unmarshal(&Conf)
	if err != nil {
		panic(fmt.Errorf("fatal error unmarshal config: %w", err))
	}
}
