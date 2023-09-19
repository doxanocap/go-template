package config

import (
	"app/internal/model"
	"github.com/doxanocap/pkg/env"
	"github.com/doxanocap/pkg/lg"
)

func InitConfig() *model.Config {
	config := model.Config{}
	err := env.LoadFile(".env")
	if err != nil {
		lg.Fatalf("config: %s", err)
	}

	err = env.Unmarshal(&config)
	if err != nil {
		lg.Fatalf("config: %s", err)
	}

	return &config
}
