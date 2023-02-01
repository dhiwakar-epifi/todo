package model

import (
	pb "example.com/grpc-todo/proto"
)

type Todo struct {
	Id          string `gorm:"primary_key"`
	Name        string
	Description string
	Done        bool
}

// TableName implementing the Tabler interface to fetch the table name .
// Done just for hygiene
func (*Todo) TableName() string {
	return "todo"
}

func NewTodo(todo *pb.Todo) *Todo {
	return &Todo{
		Id:          todo.GetId(),
		Name:        todo.GetName(),
		Description: todo.GetDescription(),
		Done:        todo.GetDone(),
	}
}

func (c *Todo) GetProto() *pb.Todo {
	todo := &pb.Todo{
		Id:          c.Id,
		Name:        c.Name,
		Description: c.Description,
		Done:        c.Done,
	}
	return todo
}
