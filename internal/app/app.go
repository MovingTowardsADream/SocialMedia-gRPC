package app

import (
	"log/slog"
	"strconv"
	grpcapp "test-gRPC/internal/app/grpc"
	"test-gRPC/internal/read_config"
	"test-gRPC/internal/service"
	"test-gRPC/internal/storage"
)

type App struct {
	GRPCSrv *grpcapp.App
}

func New(log *slog.Logger, config read_config.Config) *App {
	cfg := postgres.Config{
		Host:     config.DB.Host,
		Port:     strconv.Itoa(config.DB.Port),
		Username: config.DB.UserName,
		Password: config.DB.Password,
		DBName:   config.DB.Name,
		SSLMode:  config.DB.SSLMode,
	}

	storage, err := postgres.NewPostgresDB(cfg)
	if err != nil {
		panic(err)
	}
	authService := service.NewAuth(log, storage, config.GRPC.TokenTLL)
	twitsService := service.NewListTwit(log, storage)
	grpcApp := grpcapp.New(log, authService, twitsService, config.GRPC.Port)

	return &App{
		GRPCSrv: grpcApp,
	}
}
