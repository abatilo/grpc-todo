package fooservice

import (
	"context"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/abatilo/grpc-todo/pkg/api/v1/foo"
	"github.com/spf13/cobra"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
)

var (
	// Cmd is the exported cobra command which checks that the service is running
	Cmd = &cobra.Command{
		Use:   "fooservice",
		Short: "gRPC service",
		Run: func(cmd *cobra.Command, args []string) {
			main()
		},
	}
)

func main() {
	const port = ":8001"
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	foo.RegisterFooServiceServer(s, &foo.Server{})

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)
	defer signal.Stop(interrupt)

	g, ctx := errgroup.WithContext(ctx)

	g.Go(func() error {
		log.Printf("Starting gRPC listener on port %v", port)
		if err := s.Serve(lis); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
		return err
	})

	select {
	case <-interrupt:
		break
	case <-ctx.Done():
		break
	}

	log.Printf("Shutting down")
	s.GracefulStop()

	err = g.Wait()
	if err != nil {
		log.Fatalf("Couldn't shutdown: %v", err)
	}
}
