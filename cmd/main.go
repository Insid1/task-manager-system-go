package main

import (
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	tsmServer "go-task-manager-system"
	"go-task-manager-system/package/handler"
	"go-task-manager-system/package/repository"
	"go-task-manager-system/package/repository/postgres"
	"go-task-manager-system/package/service"
	"os"
)

func main() {
	// env
	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("error occured while getting env variables, err:%s", err.Error())
	}
	// db
	dbConfig := postgres.Config{
		Host:       os.Getenv("DB_HOST"),
		Port:       os.Getenv("DB_PORT"),
		DbName:     os.Getenv("DB_NAME"),
		DbUser:     os.Getenv("DB_USER"),
		DbPassword: os.Getenv("DB_PASSWORD"),
	}
	db := postgres.NewPostgresDb(&dbConfig)

	// packages
	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)
	// server
	server := new(tsmServer.Server)
	if err := server.Run(os.Getenv("PORT"), handlers.InitRoutes()); err != nil {
		logrus.Fatalf("Error occured while started http server: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("config")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
