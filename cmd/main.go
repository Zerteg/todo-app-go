package main

import (
	todo "github.com/Zerteg/todo-app-go"
	"log"
)

func main() {
	srv := new(todo.Server)
	if err := srv.Run("8000"); err != nil {
		log.Fatalf("error occured while runnung http server: %s", err.Error())
	}
}
