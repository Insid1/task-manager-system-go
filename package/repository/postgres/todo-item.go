package postgres

import (
	"database/sql"
	"fmt"
	todo "go-task-manager-system"
	"strings"
)

type TodoItemPostgres struct {
	db *sql.DB
}

func NewTodoItemPostgres(db *sql.DB) *TodoItemPostgres {
	return &TodoItemPostgres{db: db}
}

func (r *TodoItemPostgres) Create(item *todo.TodoItem, listId uint64) (uint64, error) {
	tx, err := r.db.Begin()

	if err != nil {
		return 0, err
	}

	var itemId uint64

	createItemQuery := fmt.Sprintf(`
		INSERT INTO %s (title, description, is_active) VALUES($1, $2, $3) RETURNING id;
	`, TodoItemsTable)

	itemRow := tx.QueryRow(createItemQuery, item.Title, item.Description, item.IsActive)
	if err = itemRow.Scan(&itemId); err != nil {
		tx.Rollback()
		return 0, err
	}

	addItemToListQuery := fmt.Sprintf(`
		INSERT INTO %s (item_id, list_id) VALUES($1, $2);
	`, ListsItemsTable)

	if _, err = tx.Exec(addItemToListQuery, itemId, listId); err != nil {
		tx.Rollback()
		return 0, err
	}
	if err = tx.Commit(); err != nil {
		tx.Rollback()
		return 0, err
	}

	return itemId, err
}

func (r *TodoItemPostgres) GetAll(listId uint64) (*[]todo.TodoItem, error) {

	query := fmt.Sprintf(`
		SELECT i.id, i.title, i.description, i.is_active
		FROM %s i
		INNER JOIN %s li
		ON i.id = li.item_id
		WHERE li.list_id = $1
	`, TodoItemsTable, ListsItemsTable)

	rows, err := r.db.Query(query, listId)

	defer rows.Close()

	if err != nil {
		return nil, err
	}

	var items []todo.TodoItem

	for i := 0; rows.Next(); i++ {
		var id uint64
		var title string
		var description string
		var isActive bool
		err = rows.Scan(&id, &title, &description, &isActive)

		if err != nil {
			return nil, err
		}

		items = append(items, todo.TodoItem{
			ID:          id,
			Title:       title,
			Description: description,
			IsActive:    isActive,
		})
	}
	err = rows.Err()
	return &items, err
}

func (r *TodoItemPostgres) GetById(listId, itemId uint64) (*todo.TodoItem, error) {
	query := fmt.Sprintf(`
		SELECT i.id, i.title, i.description, i.is_active
		FROM %s i
		INNER JOIN %s li
		ON i.id = li.item_id
		WHERE li.list_id = $1 AND i.id = $2
	`, TodoItemsTable, ListsItemsTable)

	var item todo.TodoItem
	row := r.db.QueryRow(query, listId, itemId)

	err := row.Scan(&item.ID, &item.Title, &item.Description, &item.IsActive)
	if err != nil {
		return nil, err
	}

	return &item, nil
}

func (r *TodoItemPostgres) Update(listId, itemId uint64, item *todo.UpdateTodoItemInput) error {
	values := make([]string, 0)
	args := make([]interface{}, 0)
	var argId uint8 = 1

	if item.Title != "" {
		title := fmt.Sprintf("title=$%d", argId)
		values = append(values, title)
		args = append(args, item.Title)
		argId++
	}
	if item.Description != "" {
		description := fmt.Sprintf("description=$%d", argId)
		values = append(values, description)
		args = append(args, item.Description)
		argId++
	}
	isActive := fmt.Sprintf("is_active=$%d", argId)
	values = append(values, isActive)
	args = append(args, item.IsActive)
	argId++

	setQuery := strings.Join(values, ", ")

	query := fmt.Sprintf(`
		UPDATE %s i 
		SET %s 
		FROM %s li
		WHERE i.id = li.item_id AND li.item_id = %d AND li.list_id = %d;
	`, TodoItemsTable, setQuery, ListsItemsTable, itemId, listId)
	res, err := r.db.Exec(query, args...)

	if affectedRows, _ := res.RowsAffected(); affectedRows == 0 {
		return fmt.Errorf("item for such list and user has not been found")
	}

	return err
}

func (r *TodoItemPostgres) Delete(listId, itemId uint64) error {
	query := fmt.Sprintf(`
		DELETE FROM %s i
		USING %s li
		WHERE i.id = li.item_id AND li.list_id = $1 AND i.id = $2
	`, TodoItemsTable, ListsItemsTable)

	res, err := r.db.Exec(query, listId, itemId)

	if affectedRows, _ := res.RowsAffected(); affectedRows == 0 {
		return fmt.Errorf("item for such list and user has not been found")
	}

	return err
}
