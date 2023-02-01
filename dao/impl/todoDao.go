package impl

import (
	"context"
	model "example.com/grpc-todo/dao/models"
	pb "example.com/grpc-todo/proto"
	"log"
)

type TodoDao struct {
	//db *gorm.DB
}

// func NewCrdbTodo(db *gorm.DB) *TodoDao {
func NewCrdbTodo() *TodoDao {
	return &TodoDao{
		//db: db,
	}
}

func (t *TodoDao) Create(ctx context.Context, todo *pb.Todo) (*pb.Todo, error) {
	todoModel := model.NewTodo(todo)
	log.Printf("Writing to DB")
	return todoModel.GetProto(), nil
}
