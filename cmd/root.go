package cmd

import (
	"app/internal/manager"
	"app/pkg/aws"
	"app/pkg/httpServer"
	"app/pkg/logger"
	"app/pkg/rabbitmq"
	"app/pkg/redis"
	"app/pkg/smtp"
	"context"
	"github.com/doxanocap/pkg/lg"
	"github.com/jackc/pgx/v4/pgxpool"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

func SetupManager(
	lc fx.Lifecycle,
	pool *pgxpool.Pool,
	mqClient *rabbitmq.MQClient,
	redisConn *redis.Conn,
	smtPConn *smtp.SMTP,
	awsServices *aws.Services,
	manager *manager.Manager,
) {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			manager.SetPool(pool)
			manager.SetMsgBroker(mqClient)
			manager.SetCacheConnection(redisConn)
			manager.SetMailer(smtPConn)
			manager.SetStorageProvider(awsServices.S3)
			return nil
		},
		OnStop: func(ctx context.Context) error {
			pool.Close()
			if err := mqClient.Ch.Close(); err != nil {
				return err
			}
			if err := redisConn.Close(); err != nil {
				return err
			}
			return nil
		},
	})
}

func RunServer(lc fx.Lifecycle, server *httpServer.Server, manager *manager.Manager) {
	lc.Append(fx.Hook{
		OnStart: func(context.Context) error {
			go func() {
				if err := server.Run(manager.Processor().REST().Handler().Engine()); err != nil {
					logger.Log.Fatal("failed to run REST: %v", zap.Error(err))
				}
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			lg.Info("Stopping server...")
			return server.Shutdown(ctx)
		},
	})
}
