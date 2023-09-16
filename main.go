package main

import (
	"app/cmd"
	"app/internal/config"
	"app/internal/manager"
	"app/pkg/banner"
	"app/pkg/httpServer"
	"app/pkg/logger"
	"app/pkg/postgres"
	"app/pkg/rabbitmq"
	"app/pkg/redis"
	"app/pkg/smtp"
	"go.uber.org/fx"
)

func main() {
	app := fx.New(
		fx.Provide(
			config.InitConfig,
			logger.InitLogger,
			postgres.InitConnection,
			rabbitmq.InitConnection,
			redis.InitConnection,
			smtp.InitConnection,
			manager.InitManager,
		),
		fx.Invoke(
			httpServer.InitServer,
			cmd.SetupManager,
			cmd.RunServer,
			banner.Default,
		),
	)

	app.Run()
	if err := app.Err(); err != nil {
		catcher.LogFatal(err)
	}
}
