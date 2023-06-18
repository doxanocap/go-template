package cmd

import (
	"app/internal/manager"
	"app/pkg/aws"
	"app/pkg/config"
	"app/pkg/httpServer"
	"app/pkg/logger"
	"app/pkg/postgres"
	"context"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

func Run() {
	config.InitConf()
	logger.Init(viper.GetBool("IS_PROD"), viper.GetBool("IS_JSON"))

	conn, err := postgres.Connect(context.Background(), viper.GetString("PG_DSN"))
	if err != nil {
		logger.Log.Fatal("failed connection to postgres:", zap.Error(err))
	}

	amazonWebServices := aws.Init()
	err = amazonWebServices.InitS3()
	if err != nil {
		logger.Log.Fatal("failed connection to AWS: ", zap.Error(err))
	}

	app := &App{
		Server:  httpServer.New(),
		Manager: manager.InitManager(),
	}

	app.Manager.SetPool(conn)
	app.Manager.SetStorageProvider(amazonWebServices.S3)

	// run redis and other low level services without which app should not launch
	app.Init()
}
