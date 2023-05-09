package postgres

import (
	"database/sql"
	"fmt"

	todo "go-task-manager-system"
)

type AuthPostgres struct {
	db *sql.DB
}

func NewAuthPostgres(db *sql.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (r *AuthPostgres) CreateUser(user *todo.User) (uint64, error) {
	var id uint64
	query := fmt.Sprintf("INSERT INTO %s (name, email, password_hash) values ($1, $2, $3) RETURNING id;", usersTable)

	row := r.db.QueryRow(query, user.Name, user.Email, user.Password)

	if err := row.Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}

func (r *AuthPostgres) GetUser(email, password string) (*todo.User, error) {
	var user todo.User

	query := fmt.Sprintf(`
		SELECT id, name, email, password_hash
		FROM %s
		WHERE email=$1 AND password_hash=$2;
`, usersTable)
	row := r.db.QueryRow(query, email, password)
	if err := row.Scan(&user.ID, &user.Name, &user.Email, &user.Password); err != nil {
		return nil, err
	}

	return &user, nil
}
