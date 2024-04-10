package main

import (
	"log/slog"
	"os"
	"os/signal"
	"syscall"
	"test-gRPC/internal/app"
	"test-gRPC/internal/read_config"
)

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

func main() {
	cfg := read_config.OpenConfig()

	log := setupLogger(cfg.Env)

	log.Info("Application started")

	application := app.New(log, *cfg)

	go func() {
		application.GRPCSrv.MustRun()
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	log.Info("Application shutting down")
	application.GRPCSrv.Stop()
	log.Info("Application successfully closed")
}

func setupLogger(env string) *slog.Logger {
	var log *slog.Logger

	switch env {
	case envLocal:
		log = slog.New(
			slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case envDev:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	case envProd:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}),
		)
	}

	return log
}
