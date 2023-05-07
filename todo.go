package todo

import "errors"

type TodoList struct {
	ID          uint64 `json:"id"`
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
}

type UsersList struct {
	ID     uint64
	UserId uint64
	ListId uint64
}

type TodoItem struct {
	ID          uint64 `json:"id"`
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
	IsActive    bool   `json:"isActive"`
}

type ListsItem struct {
	ID     uint64
	ListId uint64
	ItemId uint64
}

type UpdateTodoListInput struct {
	Title       string `json:"title" `
	Description string `json:"description" `
}

type UpdateTodoItemInput struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	IsActive    bool   `json:"isActive"`
}

func (i *UpdateTodoListInput) Validate() error {
	if i.Title == "" && i.Description == "" {
		return errors.New("no fields provided")
	}
	return nil
}
