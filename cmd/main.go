package main

import (
	"os"

	"github.com/Medvedevsky/todo-app"
	"github.com/Medvedevsky/todo-app/pkg/handler"
	"github.com/Medvedevsky/todo-app/pkg/repository"
	"github.com/Medvedevsky/todo-app/pkg/service"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))
	if err := initConfig(); err != nil {
		logrus.Fatalf("ошибка при чтении конфига %s", err.Error())
	}

	// грузим переменные окружения
	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("ошибка при чтении переменных окружения %s", err.Error())
	}

	db, err := repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
		Password: os.Getenv("DB.PASSWORD"),
	})

	if err != nil {
		logrus.Fatalf("ошибка при инициализации бд: %s", err.Error())
	}
	rep := repository.NewRepository(db)
	service := service.NewService(rep)
	handler := handler.NewHandler(service)
	srv := new(todo.Server)
	if err := srv.Run(viper.GetString("port"), handler.InitRoutes()); err != nil {
		logrus.Fatalf(err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
