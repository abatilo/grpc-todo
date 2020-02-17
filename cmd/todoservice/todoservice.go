package todoservice

import (
	"context"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/abatilo/grpc-todo/pkg/api/v1/todo"
	"github.com/spf13/cobra"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	healthpb "google.golang.org/grpc/health/grpc_health_v1"
)

var (
	// Cmd is the exported cobra command which checks that the service is running
	Cmd = &cobra.Command{
		Use:   "service",
		Short: "gRPC service for managing a todo list",
		Run: func(cmd *cobra.Command, args []string) {
			main()
		},
	}
)

func main() {
	const servicePort = ":8080"
	const healthcheckPort = ":8081"

	ctx, cancel := context.WithCancel(context.Background())

	// Setup
	g, ctx := errgroup.WithContext(ctx)
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)
	defer signal.Stop(interrupt)

	// Healthcheck server
	var grpcHealthServer *grpc.Server
	healthServer := health.NewServer()
	g.Go(func() error {
		grpcHealthServer = grpc.NewServer()
		healthpb.RegisterHealthServer(grpcHealthServer, healthServer)

		lis, err := net.Listen("tcp", healthcheckPort)
		if err != nil {
			log.Fatalf("failed to start healthcheck server: %v", err)
		}
		if err := grpcHealthServer.Serve(lis); err != nil {
			log.Printf("failed to shutdown healthcheck server: %v", err)
		}
		return err
	})

	// Service server
	var grpcServer *grpc.Server
	g.Go(func() error {
		grpcServer = grpc.NewServer()
		todo.RegisterTodoServiceServer(grpcServer, &todo.Server{})

		lis, err := net.Listen("tcp", servicePort)
		if err != nil {
			log.Fatalf("failed to listen: %v", err)
		}
		log.Printf("Listening to grpcServer on %v", servicePort)
		if err := grpcServer.Serve(lis); err != nil {
			log.Printf("failed to shutdown server: %v", err)
		}
		return err
	})

	select {
	case <-interrupt:
		break
	case <-ctx.Done():
		break
	}
	cancel()

	log.Printf("Shutting down")
	grpcHealthServer.GracefulStop()
	grpcServer.GracefulStop()

	err := g.Wait()
	if err != nil {
		log.Fatalf("Couldn't shutdown: %v", err)
	}
}
