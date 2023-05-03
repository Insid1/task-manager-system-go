package postgres

import (
	"database/sql"
	"fmt"
	todo "go-task-manager-system"
)

type TodoListPostgres struct {
	db *sql.DB
}

func NewTodoListPostgres(db *sql.DB) *TodoListPostgres {
	return &TodoListPostgres{db: db}
}

func (r *TodoListPostgres) Create(list todo.TodoList, userId uint64) (uint64, error) {
	tx, err := r.db.Begin()

	if err != nil {
		return 0, err
	}
	var listId uint64

	createListQuery := fmt.Sprintf("INSERT INTO %s (title, description) VALUES($1, $2) RETURNING id;", todoListsTable)
	todoListsRow := tx.QueryRow(createListQuery, list.Title, list.Description)
	err = todoListsRow.Scan(&listId)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	linkListWithUserQuery := fmt.Sprintf("INSERT INTO %s (user_id, list_id) VALUES($1, $2);", usersListsTable)
	_, err = tx.Exec(linkListWithUserQuery, userId, listId)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	return listId, tx.Commit()
}

func (r *TodoListPostgres) GetAll(userId uint64) ([]todo.TodoList, error) {
	var lists []todo.TodoList

	query := fmt.Sprintf(`
		SELECT tl.id, tl.title, tl.description
		FROM %s tl
		INNER JOIN %s ul
		ON tl.id = ul.list_id
		WHERE ul.user_id=$1;
`, todoListsTable, usersListsTable)

	rows, err := r.db.Query(query, userId)
	if err != nil {
		return lists, err
	}

	defer rows.Close()

	for i := 0; rows.Next(); i++ {
		var id uint64
		var title string
		var description string
		err = rows.Scan(&id, &title, &description)
		if err != nil {
			return lists, err
		}
		lists = append(lists, todo.TodoList{
			ID:          id,
			Title:       title,
			Description: description,
		})
	}
	err = rows.Err()
	return lists, err
}

func (r *TodoListPostgres) GetById(userId uint64, listId uint64) (todo.TodoList, error) {
	var list todo.TodoList

	query := fmt.Sprintf(`
		SELECT tl.id, tl.title, tl.description
		FROM %s tl
		INNER JOIN %s ul
		ON tl.id = ul.list_id
		WHERE ul.user_id=$1 AND tl.id=$2;
`, todoListsTable, usersListsTable)

	row := r.db.QueryRow(query, userId, listId)
	err := row.Scan(&list.ID, &list.Title, &list.Description)
	return list, err
}
