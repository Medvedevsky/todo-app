package main

import (
	"log"

	"github.com/Medvedevsky/todo-app"
	"github.com/Medvedevsky/todo-app/pkg/handler"
	"github.com/Medvedevsky/todo-app/pkg/repository"
	"github.com/Medvedevsky/todo-app/pkg/service"
	"github.com/spf13/viper"
)

func main() {
	if err := initConfig(); err != nil {
		log.Fatalf("ошибка при чтении конфига %s", err.Error())
	}
	rep := repository.NewRepository()
	service := service.NewService(rep)
	handler := handler.NewHandler(service)
	srv := new(todo.Server)
	if err := srv.Run(viper.GetString("port"), handler.InitRoutes()); err != nil {
		log.Fatalf(err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
