package main

import (
	"database/sql"
	"fmt"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	tsmServer "go-task-manager-system"
	"go-task-manager-system/package/handler"
	"go-task-manager-system/package/repository"
	"go-task-manager-system/package/service"
	"os"
)

func main() {
	// env
	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("error occured while getting env variables, err:%s", err.Error())
	}
	// db
	var (
		host       = os.Getenv("DB_HOST")
		port       = os.Getenv("DB_PORT")
		dbName     = os.Getenv("DB_NAME")
		dbUser     = os.Getenv("DB_USER")
		dbPassword = os.Getenv("DB_PASSWORD")
	)
	dataSourceName := fmt.Sprintf("host=%s port=%s dbname=%s user=%s  password=%s sslmode=disable", host, port, dbName, dbUser, dbPassword)
	fmt.Println(dataSourceName)
	db, err := sql.Open("postgres", dataSourceName)
	if err != nil {
		logrus.Fatalf("error occured while connecting database: %s", err.Error())
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		panic(err)
	}
	fmt.Println("Successfully connected to database!")

	// packages
	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)
	// server
	server := new(tsmServer.Server)
	if err := server.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		logrus.Fatalf("Error occured while started http server: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("config")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
