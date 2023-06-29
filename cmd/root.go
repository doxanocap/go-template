package cmd

import (
	"app/internal/manager"
	"app/pkg/httpServer"
	"app/pkg/logger"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

type App struct {
	Server  *httpServer.Server
	Manager *manager.Manager

	pool *pgxpool.Pool
}

func Run() {
	InitConfig()
	logger.Init(viper.GetString("ENV_MODE"), viper.GetBool("ZAP_JSON"))

	app := &App{
		Server:  httpServer.New(),
		Manager: manager.InitManager(),
	}

	app.ConnectToPostgres()
	app.ConnectToAWS()
	app.ConnectToRabbitMQ()
	app.ConnectToRedis()

	if err := app.Server.Run(app.Manager.Processor().REST().Handler().Engine()); err != nil {
		logger.Log.Fatal("failed to run REST: %v", zap.Error(err))
	}
}
