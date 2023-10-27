package app

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/gin"
	"github.com/plusik10/anti-bruteforce/config"
	v1 "github.com/plusik10/anti-bruteforce/internal/controller/http/v1"
	"github.com/plusik10/anti-bruteforce/internal/usecase"
	"github.com/plusik10/anti-bruteforce/internal/usecase/repo"
	"github.com/plusik10/anti-bruteforce/pkg/httpserver"
	"github.com/plusik10/anti-bruteforce/pkg/logger"
	"github.com/plusik10/anti-bruteforce/pkg/postgres"
)

func Run(cfg *config.Config) {
	// TODO: Init usecase -> NewRouter

	l := logger.New(cfg.Log.Level)

	pg, err := postgres.New("postgres://postgres:qwerty123@localhost:5432/anti-bruteforce")
	if err != nil {
		l.Fatal(err, " postgres is not init")
	}
	repository := repo.NewPostgres(pg)
	NetUsecase := usecase.NewNetManagerUsecase(repository)
	handler := gin.New()
	v1.NewRouter(handler, NetUsecase, l)
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
