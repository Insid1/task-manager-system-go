package postgres

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
)

const (
	usersTable      = "users"
	todoListsTable  = "todo_lists"
	usersListsTable = "users_lists"
	TodoItemsTable  = "todo_items"
	ListsItemsTable = "lists_items"
)

type Config struct {
	Host       string
	Port       string
	DbName     string
	DbUser     string
	DbPassword string
}

func NewPostgresDb(c *Config) *sql.DB {
	dataSourceName := fmt.Sprintf("host=%s port=%s dbname=%s user=%s  password=%s sslmode=disable", c.Host, c.Port, c.DbName, c.DbUser, c.DbPassword)
	db, err := sql.Open("postgres", dataSourceName)
	if err != nil {
		logrus.Fatalf("error occured while connecting database: %s", err.Error())
	}

	if err = db.Ping(); err != nil {
		panic(err)
	}
	fmt.Println("Successfully connected to postgres database!")
	return db
}
