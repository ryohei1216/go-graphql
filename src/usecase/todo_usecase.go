package usecase

import (
	"docker-go/graph/model"
	"docker-go/repository"
)

type TodoUsecase interface {
	GetTodos(userId string) ([]*model.Todo, error)
	CreateTodo(input model.NewTodo) (*model.Todo, error)
}

type todoUsecase struct {
	todoRepository repository.TodoRepository
}

func NewTodoUsecase(tr repository.TodoRepository) TodoUsecase {
	return &todoUsecase{todoRepository: tr}
}

func (tu *todoUsecase) GetTodos(userId string) ([]*model.Todo, error) {
	todos, err := tu.todoRepository.GetTodos(userId)
	if err != nil {
		return nil, err
	}
	return todos, nil
}

func (tu *todoUsecase) CreateTodo(input model.NewTodo) (*model.Todo, error) {
	todo, err := tu.todoRepository.CreateTodo(input.Text, input.UserID)
	if err != nil {
		return nil, err
	}
	return todo, nil
}