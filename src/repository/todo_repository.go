package repository

import "docker-go/graph/model"

type TodoRepository interface {
	GetTodos(userId string) ([]*model.Todo, error)
	CreateTodo(text string, userId string) (*model.Todo, error)
}