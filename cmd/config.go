package cmd

import (
	"fmt"
	"github.com/spf13/viper"
)

func InitConfig() {
	viper.SetDefault("ENV_MODE", "development")

	viper.SetDefault("ZAP_JSON", false)

	viper.SetDefault("TOKEN_MAX_AGE", 30*24*60*60*1000)
	viper.SetDefault("TOKEN_PATH", "/")
	viper.SetDefault("TOKEN_DOMAIN", "localhost")

	viper.SetDefault("JWT_REFRESH_SECRET_KEY", "secret1")
	viper.SetDefault("JWT_ACCESS_SECRET_KEY", "secret1")

	viper.SetDefault("RABBIT_MQ_DSN", "amqp://guest:guest@localhost:5672")

	//viper.SetDefault("REDIS_HOST")

	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AddConfigPath("./config")

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}

	viper.AutomaticEnv()
}
