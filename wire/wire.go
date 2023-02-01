//go:build wireinject
// +build wireinject

package wire

import (
	"example.com/grpc-todo/dao"
	services "example.com/grpc-todo/services"
	"github.com/google/wire"
)

func InitialiseTodoService() *services.TodoService {
	wire.Build(
		dao.TodoWireSet,
		services.NewTodoService,
	)
	return &services.TodoService{}
}
