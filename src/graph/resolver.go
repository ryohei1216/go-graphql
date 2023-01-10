package graph

import (
	"docker-go/usecase"
)

//go:generate go run github.com/99designs/gqlgen generate

// ここにrepositoryやservice, usecaseのstructを持つような実装になる
type Resolver struct{
	UserUsacase usecase.UserUsecase
	TodoUsecase usecase.TodoUsecase
}
