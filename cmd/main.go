package main

import (
	todo "github.com/Zerteg/todo-app-go"
	"github.com/Zerteg/todo-app-go/pkg/handler"
	"log"
)

func main() {
	handlers := new(handler.Handler)

	srv := new(todo.Server)
	if err := srv.Run("8000", handlers.IntRoutes()); err != nil {
		log.Fatalf("error occured while runnung http server: %s", err.Error())
	}
}
