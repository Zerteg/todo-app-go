package main

//sdf
import (
	todo "github.com/Zerteg/todo-app-go"
	"github.com/Zerteg/todo-app-go/pkg/handler"
	"github.com/Zerteg/todo-app-go/pkg/repository"
	"github.com/Zerteg/todo-app-go/pkg/service"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
)

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))
	if err := initConfig(); err != nil {
		logrus.Fatalf("error initializing configs: %s", err.Error())
	}
	if err := godotenv.Load(); err != nil {
		logrus.Fatalf("error loading env variable: %s", err.Error())
	}

	db, err := repository.NewPostgresDB(repository.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
		Password: os.Getenv("DB_PASSWORD"),
	})
	if err != nil {
		logrus.Fatalf("Failed on initialize db: %s", err.Error())
	}

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(todo.Server)
	if err := srv.Run(viper.GetString("port"), handlers.IntRoutes()); err != nil {
		logrus.Fatalf("error occured while runnung http server: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("C:\\Users\\Лёша\\GolandProjects\\todo-app-go\\configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
