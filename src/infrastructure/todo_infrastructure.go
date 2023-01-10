package infrastructure

import (
	"docker-go/graph/model"
	"docker-go/repository"
	"docker-go/util"
	"fmt"

	"gorm.io/gorm"
)

type todoInfrastructure struct {
	db *gorm.DB
}

func NewTodoInfrasstructure(db *gorm.DB) repository.TodoRepository {
	return &todoInfrastructure{db}
}

func (ti *todoInfrastructure) GetTodos(userId string) ([]*model.Todo, error) {
	var todos []*model.Todo
	result := ti.db.Where("user_id = ?", userId).Find(&todos)
	if result.Error != nil {
		return nil, fmt.Errorf("%v", result.Error)
	}
	return todos, nil
}

func (ti *todoInfrastructure) CreateTodo(text string, userId string) (*model.Todo, error) {
	random, _ := util.MakeRandomStr(10)
	todo := &model.Todo{ID: random, Text: text, UserID: userId}

	result := ti.db.Create(todo)
	if result.Error != nil {
		return nil, fmt.Errorf("%v", result.Error)
	}
	return todo, nil
}