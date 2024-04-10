package app

import (
	"log/slog"
	grpcapp "test-gRPC/internal/app/grpc"
	"test-gRPC/internal/service"
	"test-gRPC/internal/storage"
	"time"
)

type App struct {
	GRPCSrv *grpcapp.App
}

func New(log *slog.Logger, grpcPort int, tokenTLL time.Duration) *App {
	cfg := postgres.Config{
		Host:     "localhost",
		Port:     "5432",
		Username: "postgres",
		Password: "qwerty",
		DBName:   "postgres",
		SSLMode:  "disable",
	}
	storage, _ := postgres.NewPostgresDB(cfg)
	authService := service.NewAuth(log, storage, tokenTLL)
	twitsService := service.NewListTwit(log, storage)
	grpcApp := grpcapp.New(log, authService, twitsService, grpcPort)

	return &App{
		GRPCSrv: grpcApp,
	}
}
