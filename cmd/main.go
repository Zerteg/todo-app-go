package main

import (
	todo "github.com/Zerteg/todo-app-go"
	"github.com/Zerteg/todo-app-go/pkg/handler"
	"github.com/Zerteg/todo-app-go/pkg/repository"
	"github.com/Zerteg/todo-app-go/pkg/service"
	"github.com/spf13/viper"
	"log"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("error initializing configs: %s", err.Error())
	}

	repos := repository.NewRepository()
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(todo.Server)
	if err := srv.Run(viper.GetString("8000"), handlers.IntRoutes()); err != nil {
		log.Fatalf("error occured while runnung http server: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
