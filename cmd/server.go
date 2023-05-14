package cmd

import (
	"app/internal/manager"
	"app/pkg/httpServer"
	"app/pkg/logger"
	"github.com/jackc/pgx/v4/pgxpool"
	"go.uber.org/zap"
	"net/http"
)

type App struct {
	Server  *http.Server
	Manager *manager.Manager
}

func InitApp(conn *pgxpool.Pool) *App {
	app := &App{}
	server := httpServer.New()
	app.Manager = manager.InitManager(conn)
	if err := server.Run(app.Manager.Processor().REST().Handler().Engine); err != nil {
		logger.Log.Fatal("unable to launch RestAPI: %v", zap.Error(err))
	}

	return app
}
