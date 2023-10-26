package app

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/gin"
	"github.com/plusik10/anti-bruteforce/config"
	v1 "github.com/plusik10/anti-bruteforce/internal/controller/http/v1"
	"github.com/plusik10/anti-bruteforce/pkg/httpserver"
	"github.com/plusik10/anti-bruteforce/pkg/logger"
)

func Run(cfg *config.Config) {
	// TODO: Init usecase -> NewRouter
	handler := gin.New()
	l := logger.New(cfg.Log.Level)
	v1.NewRouter(handler, l)
	httpServer := httpserver.New(cfg, handler)

	// Waiting signal
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case s := <-interrupt:
		l.Info("app - Run - signal: " + s.String())
	case err := <-httpServer.Notiify():
		l.Error(fmt.Errorf("app - Run - httpServer.Notify: %w", err))

		// Shutdown
		err = httpServer.Shutdown()
		if err != nil {
			l.Error(fmt.Errorf("app - Run - httpServer.Shutdown: %w", err))
		}
	}
}
