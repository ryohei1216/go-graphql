package repository

import (
	"docker-go/graph/model"
)

type UserRepository interface {
	GetUsers() ([]*model.User, error)
	GetUser(id string) (*model.User, error)
}