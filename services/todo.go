package crudTodo

import (
	"context"
	pb "example.com/grpc-todo/proto"
	"github.com/google/uuid"
	"log"
)

type TodoService struct {
	pb.UnimplementedTodoServiceServer
}

func NewTodoService() *TodoService {
	return &TodoService{}
}

func (s *TodoService) CreateTodo(ctx context.Context, in *pb.NewTodo) (*pb.Todo, error) {
	log.Printf("Received todo: %v", in.GetName())
	todo := &pb.Todo{
		Name:        in.GetName(),
		Description: in.GetDescription(),
		Done:        false,
		Id:          uuid.New().String(),
	}

	return todo, nil
}
