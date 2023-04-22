package todo

type TodoList struct {
	ID          uint64 `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

type UsersList struct {
	ID     uint64
	UserId uint64
	ListId uint64
}
type TodoItem struct {
	ID          uint64 `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	IsActive    bool   `json:"isActive"`
}
type ListsItem struct {
	ID     uint64
	ListId uint64
	ItemId uint64
}