package cmd

import (
	"app/pkg/aws"
	"app/pkg/logger"
	"app/pkg/smtp"
	"go.uber.org/zap"
)

func (app *App) ConnectToAWS() {
	amazonWebServices := aws.Init()
	if err := amazonWebServices.InitS3(); err != nil {
		logger.Log.Fatal("failed connection to AWS: ", zap.Error(err))
	}

	app.Manager.SetStorageProvider(amazonWebServices.S3)
}

func (app *App) ConnectToSMTP() {
	smtpConn := smtp.Connect()
	app.Manager.SetMailer(smtpConn)
}
