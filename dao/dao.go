package dao

import (
	"context"
	"example.com/grpc-todo/dao/impl"
	pb "example.com/grpc-todo/proto"
	"github.com/google/wire"
)

var TodoWireSet = wire.NewSet(impl.NewCrdbTodo, wire.Bind(new(TodoDao), new(*impl.TodoDao)))

type TodoDao interface {
	Create(ctx context.Context, todo *pb.Todo) (*pb.Todo, error)
}
