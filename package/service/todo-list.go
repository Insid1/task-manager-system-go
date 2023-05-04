package service

import (
	todo "go-task-manager-system"
	"go-task-manager-system/package/repository"
)

type TodoListService struct {
	repo repository.TodoList
}

func NewTodoListService(repo repository.TodoList) *TodoListService {
	return &TodoListService{repo: repo}
}

func (s *TodoListService) Create(list todo.TodoList, userId uint64) (uint64, error) {
	return s.repo.Create(list, userId)
}

func (s *TodoListService) GetAll(userId uint64) ([]todo.TodoList, error) {
	return s.repo.GetAll(userId)
}

func (s *TodoListService) GetById(userId uint64, listId uint64) (todo.TodoList, error) {
	return s.repo.GetById(userId, listId)
}

func (s *TodoListService) Update(userId, listId uint64, todoList todo.UpdateTodoListInput) (todo.TodoList, error) {
	return s.repo.Update(userId, listId, todoList)
}

func (s *TodoListService) Delete(userId, listId uint64) error {
	return s.repo.Delete(userId, listId)
}
