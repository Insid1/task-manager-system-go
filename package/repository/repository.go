package repository

import (
	"database/sql"
	todo "go-task-manager-system"
	"go-task-manager-system/package/repository/postgres"
)

type Authorization interface {
	CreateUser(user todo.User) (uint64, error)
	GetUser(username, password string) (todo.User, error)
}

type TodoList interface {
	Create(list todo.TodoList, userId uint64) (uint64, error)
	GetAll(userId uint64) ([]todo.TodoList, error)
	GetById(userId uint64, listId uint64) (todo.TodoList, error)
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
		TodoList:      postgres.NewTodoListPostgres(db),
	}
}
