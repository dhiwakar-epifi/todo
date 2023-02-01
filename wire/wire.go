package wire

import (
	services "example.com/grpc-todo/services"
	"github.com/google/wire"
)

func InitialiseTodoService() *services.TodoService {
	wire.Build(
		services.NewTodoService,
	)
	return &services.TodoService{}
}
