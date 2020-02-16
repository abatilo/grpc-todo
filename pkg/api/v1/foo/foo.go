package foo

import (
	context "context"
	fmt "fmt"
	"time"

	"google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// Server is a server
type Server struct {
}

// GetFoo gets a foo
func (s *Server) GetFoo(ctx context.Context, req *FooRequest) (*FooResponse, error) {
	fmt.Println("before")
	<-time.After(time.Second * 10)
	fmt.Println("after")
	select {
	case <-ctx.Done():
		return nil, status.Errorf(codes.DeadlineExceeded, "Took too long")
	default:
		return &FooResponse{Msg: "bar"}, nil
	}
}
