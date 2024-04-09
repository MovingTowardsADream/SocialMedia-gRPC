package app

import (
	"log/slog"
	grpcapp "test-gRPC/internal/app/grpc"
	"time"
)

type App struct {
	GRPCSrv *grpcapp.App
}

func New(log *slog.Logger, grpcPort int, tokenTLL time.Duration) *App {

	grpcApp := grpcapp.New(log, grpcPort)

	return &App{
		GRPCSrv: grpcApp,
	}
}
