package service

import (
	"github.com/Medvedevsky/todo-app"
	"github.com/Medvedevsky/todo-app/pkg/repository"
)

type TodoItemService struct {
	rep repository.TodoItem
}

func NewTodoItemService(rep repository.TodoItem) *TodoItemService {
	return &TodoItemService{
		rep: rep,
	}
}

func (s *TodoItemService) Create(userId, listId int, item todo.TodoItem) (int, error) {
	return 0, nil
}

func (s *TodoItemService) GetAll(userId, listId int) ([]todo.TodoItem, error) {
	return nil, nil
}

func (s *TodoItemService) GetById(userId, itemId int) (todo.TodoItem, error) {
	return todo.TodoItem{}, nil
}

func (s *TodoItemService) Delete(userId, itemId int) error {
	return nil
}
