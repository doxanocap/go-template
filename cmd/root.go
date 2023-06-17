package cmd

import (
	"app/pkg/config"
	"app/pkg/logger"
	"app/pkg/postgres"
	"context"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

func Run() {
	config.InitConf()
	logger.Init(viper.GetBool("isPROD"), viper.GetBool("isJSON"))

	conn, err := postgres.Connect(context.Background(), viper.GetString("pg_dsn"))
	if err != nil {
		logger.Log.Fatal("failed connection to postgres:", zap.Error(err))
	}

	_ = InitApp(conn)
}
