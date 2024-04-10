package grpcapp

import (
	"fmt"
	"google.golang.org/grpc"
	"log/slog"
	"net"
	"test-gRPC/internal/grpc/authorization"
	"test-gRPC/internal/grpc/twits"
)

type App struct {
	log        *slog.Logger
	gRPCServer *grpc.Server
	port       int
}

func New(log *slog.Logger, port int) *App {
	gRPCServer := grpc.NewServer()

	authorization.Register(gRPCServer)
	twits.TwitList(gRPCServer)

	return &App{
		log:        log,
		gRPCServer: gRPCServer,
		port:       port,
	}
}

func (a *App) MustRun() {
	if err := a.Run(); err != nil {
		panic(err)
	}
}

func (a *App) Run() error {

	a.log.Info("Starting gRPC server")

	l, err := net.Listen("tcp", fmt.Sprintf(":%d", a.port))
	if err != nil {
		return fmt.Errorf(err.Error())
	}

	a.log.Info("gRPC server is running")

	if err := a.gRPCServer.Serve(l); err != nil {
		return fmt.Errorf(err.Error())
	}

	return nil
}

func (a *App) Stop() {
	a.log.Info("Stopping gRPC server")

	a.gRPCServer.GracefulStop()
}
