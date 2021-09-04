package service

import (
	"GoRestAPI/pkg/dao"
	"GoRestAPI/pkg/model"
)

type Authorization interface {
	CreateUser(user model.User) (int, error)
}

type TodoList interface {
}

type TodoItem interface {
}

type Service struct {
	Authorization
	TodoList
	TodoItem
}

func NewService(repo *dao.Dao) *Service {
	return &Service{
		Authorization: NewAuthService(repo.Authorization),
	}
}
