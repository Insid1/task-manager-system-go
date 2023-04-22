package main

import (
	tsmServer "go-task-manager-system"
	"go-task-manager-system/package/handler"
	"log"
)

const PORT = "9090"

func main() {
	handlers := new(handler.Handler) //todo:  check why new
	server := new(tsmServer.Server)
	if err := server.Run(PORT, handlers.InitRoutes()); err != nil {
		log.Fatalf("Error occured while started http server: %s", err.Error())
	}
}
