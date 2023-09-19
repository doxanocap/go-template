package httpServer

import (
	"app/internal/model"
	"context"
	"fmt"
	"net/http"
	"time"
)

type Server struct {
	cfg *model.Config

	httpServer *http.Server
}

func InitServer(cfg *model.Config) *Server {
	return &Server{
		cfg: cfg,
	}
}

func (s *Server) Run(handler http.Handler) error {
	s.httpServer = &http.Server{
		Addr:           ":" + s.cfg.ServerPort,
		Handler:        handler,
		MaxHeaderBytes: 1 << 20, // 1mb
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
	}
	fmt.Printf("Listening on port: %s\n\n", s.cfg.ServerPort)
	return s.httpServer.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
