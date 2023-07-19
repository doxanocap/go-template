package cmd

import (
	"app/pkg/aws"
	"app/pkg/logger"
	"app/pkg/postgres"
	"app/pkg/rabbitmq"
	"app/pkg/redis"
	"app/pkg/smtp"
	"context"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

func (app *App) ConnectToPostgres() {
	conn, err := postgres.Connect(viper.GetString("PG_DSN"))
	if err != nil {
		logger.Log.Fatal("failed connection to postgres:", zap.Error(err))
	}
	app.Manager.SetConn(conn)
}

func (app *App) ConnectToAWS() {
	amazonWebServices := aws.Init()
	if err := amazonWebServices.InitS3(); err != nil {
		logger.Log.Fatal("failed connection to AWS: ", zap.Error(err))
	}

	app.Manager.SetStorageProvider(amazonWebServices.S3)
}

func (app *App) ConnectToRabbitMQ() {
	msgBroker, err := rabbitmq.Connect()
	if err != nil {
		logger.Log.Fatal("failed connection to RabbitMQ: ", zap.Error(err))
	}
	app.Manager.SetMsgBroker(msgBroker)
}

func (app *App) ConnectToRedis() {
	redisConn, err := redis.Connect(context.Background(), redis.Config{
		Host:      viper.GetString("REDIS_HOST"),
		Password:  viper.GetString("REDIS_PASSWORD"),
		DB:        viper.GetInt("REDIS_DB"),
		KeyPrefix: viper.GetString("REDIS_PREFIX"),
	})
	if err != nil {
		logger.Log.Fatal("failed connection to Redis: ", zap.Error(err))
	}
	app.Manager.SetCacheConnection(redisConn)
}

func (app *App) ConnectToSMTP() {
	smtpConn := smtp.Connect()
	app.Manager.SetMailer(smtpConn)
}
