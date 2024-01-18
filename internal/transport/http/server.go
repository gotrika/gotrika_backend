package http

import (
	"context"
	"net/http"

	"github.com/gotrika/gotrika_backend/internal/config"
)

type Server struct {
	httpServer *http.Server
}

func NewHTTPServer(cfg *config.Config, handler http.Handler) *Server {
	return &Server{
		httpServer: &http.Server{
			Addr:    ":" + cfg.HTTPConfig.Port,
			Handler: handler,
		},
	}
}

func (s *Server) Run() error {
	return s.httpServer.ListenAndServe()
}

func (s *Server) Stop(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
