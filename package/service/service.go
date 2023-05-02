package service

import (
	todo "go-task-manager-system"
	"go-task-manager-system/package/repository"
)

type Authorization interface {
	CreateUser(user todo.User) (uint64, error)
	GenerateToken(username, password string) (string, error)
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

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: newAuthService(repos.Authorization),
	}
}
