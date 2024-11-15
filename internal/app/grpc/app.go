package grpcapp

import (
	"fmt"
	"github.com/EvgeniyMdr/commentService/internal/config"
	"github.com/EvgeniyMdr/commentService/internal/grpc/server"
	"github.com/EvgeniyMdr/commentService/internal/services"
	"google.golang.org/grpc"
	"net"
)

type App struct {
	gRPCServer *grpc.Server
}

func New(commentService services.Service) *App {
	gRPCServer := grpc.NewServer()

	server.Register(gRPCServer, commentService)

	return &App{
		gRPCServer: gRPCServer,
	}
}

func (a *App) Run() error {
	mainConfig := config.NewServiceConfig()
	grpcSettings := mainConfig.GetGRPCSettings()

	fmt.Printf("grpcSettings %v", grpcSettings)

	address := fmt.Sprintf("%s:%s", grpcSettings.Host, grpcSettings.Port)

	fmt.Printf("address: %s", address)

	l, err := net.Listen("tcp", address)

	if err != nil {
		return fmt.Errorf("%w", err)
	}

	if err := a.gRPCServer.Serve(l); err != nil {
		return fmt.Errorf("%w", err)
	}

	return nil
}

func (a *App) MustRun() {
	if err := a.Run(); err != nil {
		panic(err)
	}
}

func (a *App) Stop() {
	a.gRPCServer.GracefulStop()
}
