package main

import (
	"fmt"
	"log"
	"os"

	"github.com/Medvedevsky/todo-app"
	"github.com/Medvedevsky/todo-app/pkg/handler"
	"github.com/Medvedevsky/todo-app/pkg/repository"
	"github.com/Medvedevsky/todo-app/pkg/service"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
)

func main() {

	if err := initConfig(); err != nil {
		log.Fatalf("ошибка при чтении конфига %s", err.Error())
	}

	// грузим переменные окружения
	if err := godotenv.Load(); err != nil {
		log.Fatalf("ошибка при чтении переменных окружения %s", err.Error())
	}

	test := viper.GetString("db.sslmode")
	test2 := os.Getenv("DB.PASSWORD")

	fmt.Println(test)
	fmt.Println(test2)

	db, err := repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
		Password: os.Getenv("DB.PASSWORD"),
	})

	if err != nil {
		log.Fatalf("ошибка при инициализации бд: %s", err.Error())
	}
	rep := repository.NewRepository(db)
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
