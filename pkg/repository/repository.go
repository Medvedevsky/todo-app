package repository

import (
	"github.com/Medvedevsky/todo-app"
	"github.com/jmoiron/sqlx"
)

type Repository struct {
	Autharization Autharization
	TodoItem      TodoItem
	TodoList      TodoList
}

type Autharization interface {
	CreateUser(user todo.User) (int, error)
	GetUser(login string, password string) (todo.User, error)
}

type TodoItem interface {
}

type TodoList interface {
	Create(userId int, list todo.TodoList) (int, error)
	GetAll(userId int) ([]todo.TodoList, error)
	GetById(userId, listId int) (todo.TodoList, error)
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Autharization: NewAuthPostgres(db),
	}
}
