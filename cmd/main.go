package main

import (
	"log"
	"todo"
	"todo/pkg/handlers"
)

func main() {
	handler := new(handlers.Handler)
	srv := new(todo.Server)
	if err := srv.Run(":8000", handler.InitRoutes()); err != nil {
		log.Fatalf("error occured while running http server: %s", err.Error())
	}
}
