package dao

import (
	"GoRestAPI/pkg/dao/postgres"
	"GoRestAPI/pkg/model"
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user model.User) (int, error)
}

type TodoList interface {
}

type TodoItem interface {
}

type Dao struct {
	Authorization
	TodoList
	TodoItem
}

func NewDao(db *sqlx.DB) *Dao {
	return &Dao{
		Authorization: postgres.NewAuthPostgres(db),
	}
}
