package main

import (
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"github.com/zlietapki/microboiler_rest_server/internal/config"
	"github.com/zlietapki/microboiler_rest_server/internal/repository"
	"github.com/zlietapki/microboiler_rest_server/internal/rest_handler"
	"github.com/zlietapki/microboiler_rest_server/internal/usecase"
	"github.com/zlietapki/microboiler_rest_server/pkg/httpserver"
)

func main() {
	cfg := config.New()

	repo := repository.New()
	uc := usecase.New(repo)

	restHandler := rest_handler.New(uc)

	httpServer := httpserver.New(cfg.HTTPListen)

	api := httpServer.Srv.Group("/api")
	api.GET("/users", restHandler.GetUsers)
	api.POST("/users", restHandler.CreateUser)

	httpErrCh := httpServer.Start()
	slog.Info("HTTP listening", "addr", cfg.HTTPListen)

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)
	select {
	case <-signals:
	case err := <-httpErrCh:
		panic("http server" + err.Error())
	}
}
