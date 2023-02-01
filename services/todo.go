package crudTodo

import (
	"context"
	"example.com/grpc-todo/dao"
	pb "example.com/grpc-todo/proto"
	"github.com/google/uuid"
	"log"
)

type TodoService struct {
	pb.UnimplementedTodoServiceServer
	todoDao dao.TodoDao
}

func NewTodoService(todoDao dao.TodoDao) *TodoService {
	return &TodoService{
		todoDao: todoDao,
	}
}

func (s *TodoService) CreateTodo(ctx context.Context, in *pb.NewTodo) (*pb.Todo, error) {
	log.Printf("Received todo: %v", in.GetName())
	todo := &pb.Todo{
		Name:        in.GetName(),
		Description: in.GetDescription(),
		Done:        false,
		Id:          uuid.New().String(),
	}
	s.todoDao.Create(ctx, todo)

	return todo, nil
}
