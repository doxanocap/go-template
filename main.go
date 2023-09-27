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

//	@title			go-template
//	@version		1.0
//	@description	Swagger документация для go-template.
//	@termsOfService	http://swagger.io/terms/

// @host		localhost:8080
// @BasePath	/
func main() {
	app := fx.New(
		fx.NopLogger,
		fx.Provide(
			config.InitConfig,
			postgres.InitConnection,
			rabbitmq.InitConnection,
			redis.InitConnection,
			smtp.InitConnection,
			aws.InitServices,
			manager.InitManager,
			httpServer.InitServer,
		),
		fx.Invoke(
			logger.InitLogger,
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
