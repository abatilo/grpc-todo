package fooclient

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/abatilo/grpc-todo/pkg/api/v1/foo"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
)

var (
	// Cmd is the exported cobra command which checks that the service is running
	Cmd = &cobra.Command{
		Use:   "fooclient",
		Short: "gRPC client",
		Run: func(cmd *cobra.Command, args []string) {
			main()
		},
	}
)

func main() {
	conn, err := grpc.Dial("localhost:8001", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}
	defer conn.Close()
	c := foo.NewFooServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*20)
	req := &foo.FooRequest{}
	f, err := c.GetFoo(ctx, req)
	if err != nil {
		log.Fatalf("couldn't get foo: %v", err)
	}
	fmt.Printf("foo.msg: %v", f.GetMsg())
	cancel()
}
