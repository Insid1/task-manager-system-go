package service

import "go-task-manager-system/package/repository"

type TodoListService struct {
	repo repository.TodoList
}

func newTodoListService(repo repository.TodoList) *TodoListService {
	return &TodoListService{repo: repo}
}
