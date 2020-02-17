package todoclient

import (
	"context"
	"log"
	"time"

	"github.com/abatilo/grpc-todo/pkg/api/v1/todo"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
)

var (
	// Cmd is the exported cobra command which checks that the service is running
	Cmd = &cobra.Command{
		Use:   "client",
		Short: "gRPC client for fetching a todo list",
		Run: func(cmd *cobra.Command, args []string) {
			main()
		},
	}
)

func main() {
	conn, err := grpc.Dial("localhost:8080", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}
	defer conn.Close()
	c := todo.NewTodoServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*60)

	{
		req := &todo.ListTodoRequest{}
		resp, err := c.ListTodos(ctx, req)
		if err != nil {
			log.Fatalf("couldn't list todos: %v", err)
		}
		log.Printf("ListTodoResponse: %v", resp)
	}

	{
		req := &todo.AddTodoRequest{Description: "Testing"}
		resp, err := c.AddTodo(ctx, req)
		if err != nil {
			log.Fatalf("Couldn't add todo: %v", err)
		}
		log.Printf("Just added todo with ID: %v", resp.GetTodoId())
	}

	{
		req := &todo.ListTodoRequest{}
		resp, err := c.ListTodos(ctx, req)
		if err != nil {
			log.Fatalf("couldn't list todos: %v", err)
		}
		log.Printf("ListTodoResponse: %v", resp)
	}

	{
		req := &todo.UpdateTodoRequest{Complete: true}
		resp, err := c.UpdateTodo(ctx, req)
		if err != nil {
			log.Fatalf("couldn't update todo: %v", err)
		}
		log.Printf("UpdateTodoResponse: %v", resp)
	}

	cancel()
}
