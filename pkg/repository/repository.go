package repository

import "github.com/Medvedevsky/todo-app"

type Repository struct {
	Autharization Autharization
	TodoItem      TodoItem
	TodoList      TodoList
}

type Autharization interface {
	CreateUser(user todo.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
}

type TodoItem interface {
}

type TodoList interface {
	Create(userId int, list todo.TodoList) (int, error)
	GetAll(userId int) ([]todo.TodoList, error)
	GetById(userId, listId int) (todo.TodoList, error)
}

func NewRepository() *Repository {
	return &Repository{}
}
