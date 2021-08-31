package service

import "GoRestAPI/pkg/dao"

type Authorization interface {

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

func NewService(repos *dao.Dao) *Service {
	return &Service{}
}

