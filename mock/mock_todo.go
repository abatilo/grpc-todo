// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/abatilo/grpc-todo/pkg/api/v1/todo (interfaces: TodoServiceClient)

// Package mocks is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"

	todo "github.com/abatilo/grpc-todo/pkg/api/v1/todo"
	gomock "github.com/golang/mock/gomock"
	grpc "google.golang.org/grpc"
)

// MockTodoServiceClient is a mock of TodoServiceClient interface
type MockTodoServiceClient struct {
	ctrl     *gomock.Controller
	recorder *MockTodoServiceClientMockRecorder
}

// MockTodoServiceClientMockRecorder is the mock recorder for MockTodoServiceClient
type MockTodoServiceClientMockRecorder struct {
	mock *MockTodoServiceClient
}

// NewMockTodoServiceClient creates a new mock instance
func NewMockTodoServiceClient(ctrl *gomock.Controller) *MockTodoServiceClient {
	mock := &MockTodoServiceClient{ctrl: ctrl}
	mock.recorder = &MockTodoServiceClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockTodoServiceClient) EXPECT() *MockTodoServiceClientMockRecorder {
	return m.recorder
}

// AddTodo mocks base method
func (m *MockTodoServiceClient) AddTodo(arg0 context.Context, arg1 *todo.AddTodoRequest, arg2 ...grpc.CallOption) (*todo.AddTodoResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "AddTodo", varargs...)
	ret0, _ := ret[0].(*todo.AddTodoResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddTodo indicates an expected call of AddTodo
func (mr *MockTodoServiceClientMockRecorder) AddTodo(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddTodo", reflect.TypeOf((*MockTodoServiceClient)(nil).AddTodo), varargs...)
}

// GetTodo mocks base method
func (m *MockTodoServiceClient) GetTodo(arg0 context.Context, arg1 *todo.GetTodoRequest, arg2 ...grpc.CallOption) (*todo.GetTodoResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetTodo", varargs...)
	ret0, _ := ret[0].(*todo.GetTodoResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTodo indicates an expected call of GetTodo
func (mr *MockTodoServiceClientMockRecorder) GetTodo(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTodo", reflect.TypeOf((*MockTodoServiceClient)(nil).GetTodo), varargs...)
}

// ListTodos mocks base method
func (m *MockTodoServiceClient) ListTodos(arg0 context.Context, arg1 *todo.ListTodoRequest, arg2 ...grpc.CallOption) (*todo.ListTodoResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListTodos", varargs...)
	ret0, _ := ret[0].(*todo.ListTodoResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListTodos indicates an expected call of ListTodos
func (mr *MockTodoServiceClientMockRecorder) ListTodos(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListTodos", reflect.TypeOf((*MockTodoServiceClient)(nil).ListTodos), varargs...)
}

// UpdateTodo mocks base method
func (m *MockTodoServiceClient) UpdateTodo(arg0 context.Context, arg1 *todo.UpdateTodoRequest, arg2 ...grpc.CallOption) (*todo.UpdateTodoResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UpdateTodo", varargs...)
	ret0, _ := ret[0].(*todo.UpdateTodoResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateTodo indicates an expected call of UpdateTodo
func (mr *MockTodoServiceClientMockRecorder) UpdateTodo(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateTodo", reflect.TypeOf((*MockTodoServiceClient)(nil).UpdateTodo), varargs...)
}
