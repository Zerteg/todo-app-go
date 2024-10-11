package main

import (
	todo "github.com/Zerteg/todo-app-go"
	"github.com/Zerteg/todo-app-go/pkg/handler"
	"github.com/Zerteg/todo-app-go/pkg/repository"
	"github.com/Zerteg/todo-app-go/pkg/service"
	"log"
)

func main() {
	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(todo.Server)
	if err := srv.Run("8000", handlers.IntRoutes()); err != nil {
		log.Fatalf("error occured while runnung http server: %s", err.Error())
	}
}
