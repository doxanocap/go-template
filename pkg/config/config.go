package config

import (
	"fmt"
	"github.com/spf13/viper"
)

func InitConf() {
	viper.SetDefault("isProd", false)
	viper.SetDefault("isJson", false)
	viper.SetDefault("EnvMode", "development")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("../../../")

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}

	viper.AutomaticEnv()
}
