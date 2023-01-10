package usecase

import (
	"docker-go/graph/model"
	"docker-go/repository"
)

type UserUsecase interface {
	GetUsers() ([]*model.User, error)
	GetUser(id string) (*model.User, error)
}

type userUsecase struct {
	userRepository repository.UserRepository
}

func NewUserUsecase(ur repository.UserRepository) UserUsecase {
	return &userUsecase{userRepository: ur}
}

func (uu *userUsecase) GetUsers() ([]*model.User, error) {
	users, err := uu.userRepository.GetUsers()
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (uu *userUsecase) GetUser(id string) (*model.User, error) {
	user, err := uu.userRepository.GetUser(id)
	if err != nil {
		return nil, err
	}
	return user, nil
}

