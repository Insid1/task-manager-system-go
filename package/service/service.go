package service

import (
	todo "go-task-manager-system"
	"go-task-manager-system/package/repository"
)

type Authorization interface {
	CreateUser(user todo.User) (uint64, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(accessToken string) (uint64, error)
}

type TodoList interface {
	Create(list *todo.TodoList, userId uint64) (uint64, error)
	GetAll(userId uint64) (*[]todo.TodoList, error)
	GetById(userId, listId uint64) (*todo.TodoList, error)
	Update(userId, listId uint64, todoList *todo.UpdateTodoListInput) error
	Delete(userId, listId uint64) error
}

type TodoItem interface {
	Create(item *todo.TodoItem, userId, listId uint64) (uint64, error)
	GetAll(userId, listId uint64) (*[]todo.TodoItem, error)
	GetById(userId, listId, itemId uint64) (*todo.TodoItem, error)
	Delete(userId, listId, itemId uint64) error
}

type Service struct {
	Authorization
	TodoList
	TodoItem
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		TodoList:      NewTodoListService(repos.TodoList),
		TodoItem:      NewTodoItemService(repos.TodoItem, repos.TodoList),
	}
}
