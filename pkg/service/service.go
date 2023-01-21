package service

import (
	"github.com/Medvedevsky/todo-app"
	"github.com/Medvedevsky/todo-app/pkg/repository"
)

type Service struct {
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
	Create(userId, listId int, item todo.TodoItem) (int, error)
	GetAll(userId, listId int) ([]todo.TodoItem, error)
	GetById(userId, itemId int) (todo.TodoItem, error)
	Delete(userId, itemId int) error
}

type TodoList interface {
	Create(userId int, list todo.TodoList) (int, error)
	GetAll(userId int) ([]todo.TodoList, error)
	GetById(userId, listId int) (todo.TodoList, error)
}

func NewService(rep *repository.Repository) *Service {
	return &Service{
		TodoList: NewTodoListService(rep.TodoList),
		TodoItem: NewTodoItemService(rep.TodoItem),
	}
}
