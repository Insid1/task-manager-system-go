package main

import (
	"github.com/spf13/viper"
	tsmServer "go-task-manager-system"
	"go-task-manager-system/package/handler"
	"go-task-manager-system/package/repository"
	"go-task-manager-system/package/service"
	"log"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("error occured while initialize config: %s", err.Error())
	}
	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	server := new(tsmServer.Server)
	if err := server.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		log.Fatalf("Error occured while started http server: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("config")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
