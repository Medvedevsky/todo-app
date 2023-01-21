package service

import (
	"github.com/Medvedevsky/todo-app"
	"github.com/Medvedevsky/todo-app/pkg/repository"
)

type TodoListService struct {
	rep repository.TodoList
}

func NewTodoListService(rep repository.TodoList) *TodoListService {
	return &TodoListService{
		rep: rep,
	}
}

func (s *TodoListService) Create(userId int, list todo.TodoList) (int, error) {
	return 0, nil
}

func (s *TodoListService) GetAll(userId int) ([]todo.TodoList, error) {
	return nil, nil
}

func (s *TodoListService) GetById(userId, listId int) (todo.TodoList, error) {
	return todo.TodoList{}, nil
}

func (s *TodoListService) Delete(userId, listId int) error {
	return nil
}
