package main

import (
	"log"

	"github.com/Medvedevsky/todo-app"
	"github.com/Medvedevsky/todo-app/pkg/handler"
)

func main() {
	handler := new(handler.Handler)
	srv := new(todo.Server)
	if err := srv.Run("8080", handler.InitRoutes()); err != nil {
		log.Fatalf(err.Error())
	}
}
