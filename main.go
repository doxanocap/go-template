package main

import (
	"app/cmd"
	"app/internal/config"
	"app/internal/manager"
	"app/pkg/aws"
	"app/pkg/banner"
	"app/pkg/httpServer"
	"app/pkg/logger"
	"app/pkg/postgres"
	"app/pkg/rabbitmq"
	"app/pkg/redis"
	"app/pkg/smtp"
	"github.com/doxanocap/pkg/lg"
	"go.uber.org/fx"
)

func main() {
	app := fx.New(
		fx.NopLogger,
		fx.Provide(
			config.InitConfig,
			logger.InitLogger,
			postgres.InitConnection,
			rabbitmq.InitConnection,
			redis.InitConnection,
			smtp.InitConnection,
			aws.InitServices,
			manager.InitManager,
			httpServer.InitServer,
		),
		fx.Invoke(
			cmd.SetupManager,
			cmd.RunServer,
			banner.Default,
		),
	)

	app.Run()
	if err := app.Err(); err != nil {
		lg.Fatal(err)
	}
}
