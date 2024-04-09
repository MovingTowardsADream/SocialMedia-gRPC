package main

import (
	"log/slog"
	"os"
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

	application := app.New(log, cfg.GRPC.Port, cfg.GRPC.TokenTLL)

	application.GRPCSrv.MustRun()
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
