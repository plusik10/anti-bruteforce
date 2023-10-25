package httpserver

import (
	"context"
	"net/http"
	"time"

	"github.com/golangci/golangci-lint/pkg/config"
)

type Server struct {
	server          *http.Server
	notify          chan error
	shutdownTimeout time.Duration
}

func New(cfg config.Config, handler http.Handler) *Server {
	httpServer := &http.Server{
		Handler:      handler,
		ReadTimeout:  cfg.HTTP.Timeout,
		Addr:         cfg.HTTP.Port,
		WriteTimeout: cfg.HTTP.Timeout,
	}

	s := &Server{
		server:          httpServer,
		notify:          make(chan error, 1),
		shutdownTimeout: 3 * time.Second,
	}

	s.start()
	return s

}

func (s *Server) start() {
	go func() {
		s.notify <- s.server.ListenAndServe()
		close(s.notify)
	}()
}

func (s *Server) Notiify() <-chan error {
	return s.notify
}

func (s *Server) Shutdown() error {
	ctx, cancel := context.WithTimeout(context.Background(), s.shutdownTimeout)
	defer cancel()
	return s.server.Shutdown(ctx)
}
