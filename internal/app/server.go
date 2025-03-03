package app

import (
	"context"
	"net"
	"net/http"

	"github.com/ryoeuyo/goauth/internal/config"
)

type server struct {
	server *http.Server
}

func newServer(cfg config.HTTPServer, handler http.Handler) server {
	return server{
		server: &http.Server{
			Addr:           net.JoinHostPort(cfg.Host, cfg.Port),
			Handler:        handler,
			ReadTimeout:    cfg.Timeout,
			WriteTimeout:   cfg.Timeout,
			IdleTimeout:    cfg.IdleTimeout,
			MaxHeaderBytes: 1 << 20,
		},
	}
}

func (s *server) start() {
	_ = s.server.ListenAndServe()
}

func (s *server) shutdown(ctx context.Context) error {
	return s.server.Shutdown(ctx)
}
