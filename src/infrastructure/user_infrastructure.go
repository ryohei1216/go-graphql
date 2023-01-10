package infrastructure

import (
	"docker-go/graph/model"
	"docker-go/repository"
	"fmt"

	"gorm.io/gorm"
)

type UserInfrastructure struct {
	db *gorm.DB
}

func NewUserInfrasstructure(db *gorm.DB) repository.UserRepository {
	return &UserInfrastructure{db}
}

func(ui *UserInfrastructure) GetUsers() ([]*model.User, error){
	var users []*model.User
	result := ui.db.Find(&users)
	if result.Error != nil {
		return nil, fmt.Errorf("%v", result.Error)
	}
	return users, nil
}

func(ui *UserInfrastructure) GetUser(id string) (*model.User, error){
	var user *model.User
	result := ui.db.Where("id = ?", id).Find(&user)
	if result.Error != nil {
		return nil, fmt.Errorf("%v", result.Error)
	}
	return user, nil
}