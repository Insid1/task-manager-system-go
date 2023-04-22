package main

import (
	tsmServer "go-task-manager-system"
	"go-task-manager-system/package/handler"
	"go-task-manager-system/package/repository"
	"go-task-manager-system/package/service"
	"log"
)

const PORT = "9090"

func main() {
	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	server := new(tsmServer.Server)
	if err := server.Run(PORT, handlers.InitRoutes()); err != nil {
		log.Fatalf("Error occured while started http server: %s", err.Error())
	}
}
