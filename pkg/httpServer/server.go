package httpServer

import (
	"context"
	"fmt"
	"github.com/spf13/viper"
	"net/http"
	"time"
)

type Server struct {
	httpServer *http.Server
}

func New() *Server {
	return &Server{}
}

func (s *Server) Run(handler http.Handler) error {
	s.httpServer = &http.Server{
		Addr:           ":" + viper.GetString("PORT"),
		Handler:        handler,
		MaxHeaderBytes: 1 << 20, // 1mb
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
	}
	fmt.Printf("[GIN-debug]         Listening on port: %s\n\n", viper.GetString("PORT"))
	return s.httpServer.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
