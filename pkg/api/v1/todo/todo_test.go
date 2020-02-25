package todo_test

import (
	"context"
	"testing"

	"github.com/abatilo/grpc-todo/mock"
	"github.com/abatilo/grpc-todo/pkg/api/v1/todo"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

// TestAddTodo just a simple test to try mocking a gRPC client
func TestAddTodo(t *testing.T) {
	assert := assert.New(t)
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockClient := mock.NewMockTodoServiceClient(ctrl)

	mockClient.EXPECT().AddTodo(
		gomock.Any(),
		gomock.Any(),
	).Return(&todo.AddTodoResponse{TodoId: 0}, nil)

	resp, err := mockClient.AddTodo(context.Background(), &todo.AddTodoRequest{})
	assert.Nil(err)

	assert.EqualValues(&todo.AddTodoResponse{TodoId: 0}, resp)
}
