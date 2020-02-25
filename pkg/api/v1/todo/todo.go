package todo

//go:generate mockgen -package mock -destination=../../../../mock/mock_todo.go github.com/abatilo/grpc-todo/pkg/api/v1/todo TodoServiceClient

import (
	context "context"
	fmt "fmt"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Server is a server
type Server struct {
	todos []*Todo
}

var (
	// ErrTodoNotFound represents when a todo doesn't exist but is attempted to be accessed
	ErrTodoNotFound = status.Error(codes.NotFound, "This provided ID doesn't exist in the list of available Todos")
)

// AddTodo add a todo to the list
func (s *Server) AddTodo(ctx context.Context, r *AddTodoRequest) (*AddTodoResponse, error) {
	fmt.Println("AddTodo")
	s.todos = append(s.todos, &Todo{Description: r.GetDescription()})
	resp := &AddTodoResponse{TodoId: uint32(len(s.todos))}
	return resp, nil
}

// GetTodo gets a todo
func (s *Server) GetTodo(ctx context.Context, r *GetTodoRequest) (*GetTodoResponse, error) {
	fmt.Println("GetTodo")
	if uint32(len(s.todos)) < r.GetTodoId() {
		return nil, ErrTodoNotFound
	}
	todo := s.todos[r.GetTodoId()]
	resp := &GetTodoResponse{Todo: todo}
	return resp, nil
}

// ListTodos lists every todo
func (s *Server) ListTodos(ctx context.Context, r *ListTodoRequest) (*ListTodoResponse, error) {
	fmt.Println("ListTodos")
	resp := &ListTodoResponse{Todos: s.todos}
	return resp, nil
}

// UpdateTodo will update a todo
func (s *Server) UpdateTodo(ctx context.Context, r *UpdateTodoRequest) (*UpdateTodoResponse, error) {
	fmt.Println("UpdateTodo")
	if uint32(len(s.todos)) < r.GetTodoId() {
		return nil, ErrTodoNotFound
	}

	t := s.todos[r.GetTodoId()]
	t.Complete = r.GetComplete()
	resp := &UpdateTodoResponse{Todo: t}
	return resp, nil
}
