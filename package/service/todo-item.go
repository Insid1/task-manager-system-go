package service

import (
	todo "go-task-manager-system"
	"go-task-manager-system/package/repository"
)

type TodoItemService struct {
	repoItem repository.TodoItem
	listRepo repository.TodoList
}

func NewTodoItemService(repoItem repository.TodoItem, repoList repository.TodoList) *TodoItemService {
	return &TodoItemService{repoItem: repoItem, listRepo: repoList}
}

func (s *TodoItemService) Create(item *todo.TodoItem, userId uint64, listId uint64) (uint64, error) {
	_, err := s.listRepo.GetById(userId, listId)
	if err != nil {
		return 0, err
	}

	return s.repoItem.Create(item, listId)
}

func (s *TodoItemService) GetAll(userId, listId uint64) (*[]todo.TodoItem, error) {
	_, err := s.listRepo.GetById(userId, listId)
	if err != nil {
		return nil, err
	}

	return s.repoItem.GetAll(listId)

}

func (s *TodoItemService) GetById(userId, listId, itemId uint64) (*todo.TodoItem, error) {
	_, err := s.listRepo.GetById(userId, listId)
	if err != nil {
		return nil, err
	}

	item, err := s.repoItem.GetById(listId, itemId)

	if err != nil {
		return nil, err
	}

	return item, nil
}

func (s *TodoItemService) Update(userId, listId, itemId uint64, todoItem *todo.UpdateTodoItemInput) error {
	_, err := s.listRepo.GetById(userId, listId)
	if err != nil {
		return err
	}

	err = s.repoItem.Update(listId, itemId, todoItem)
	return err
}

func (s *TodoItemService) Delete(userId, listId, itemId uint64) error {
	_, err := s.listRepo.GetById(userId, listId)
	if err != nil {
		return err
	}

	err = s.repoItem.Delete(listId, itemId)
	return err

}
