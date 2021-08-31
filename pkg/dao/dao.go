package dao

import "github.com/jmoiron/sqlx"

type Authorization interface {
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
	return &Dao{}
}
