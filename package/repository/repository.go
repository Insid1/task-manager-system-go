package repository

import (
	"database/sql"
	todo "go-task-manager-system"
	"go-task-manager-system/package/repository/postgres"
)

type Authorization interface {
	CreateUser(user todo.User) (uint64, error)
}

type TodoList interface {
}

type TodoItem interface {
}

type Repository struct {
	Authorization
	TodoList
	TodoItem
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		Authorization: postgres.NewAuthPostgres(db),
	}
}
