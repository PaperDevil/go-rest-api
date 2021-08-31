package dao

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

func NewDao() *Dao {
	return &Dao{}
}