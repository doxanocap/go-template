package cmd

import (
	"app/internal/manager"
	"app/pkg/httpServer"
	"app/pkg/logger"
	"github.com/jackc/pgx/v4/pgxpool"
	"go.uber.org/zap"
)

type App struct {
	Server  *httpServer.Server
	Manager *manager.Manager
}

func InitApp(conn *pgxpool.Pool) *App {
	app := &App{
		Server:  httpServer.New(),
		Manager: manager.InitManager(conn),
	}

	if err := app.Server.Run(app.Manager.Processor().REST().Handler().Engine); err != nil {
		logger.Log.Fatal("failed to run REST: %v", zap.Error(err))
	}

	return app
}
