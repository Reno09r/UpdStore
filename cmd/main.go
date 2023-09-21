package main

import (
	"os"

	store "github.com/Reno09r/Store"
	"github.com/Reno09r/Store/pkg/handler"
	"github.com/Reno09r/Store/pkg/repository"
	"github.com/Reno09r/Store/pkg/service"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))
	if err := initConfig(); err != nil {
		logrus.Fatalf("error initializating confings: %s", err.Error())
	}
	if err := godotenv.Load("c:\\Users\\Reno\\Desktop\\My\\Projects\\Go\\UpdStore\\.env"); err != nil {
		logrus.Fatalf("error initializating confings: %s", err.Error())
	}
	db, err := repository.NewPostgresDB(repository.DBConfig{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		Password: os.Getenv("DB_PASSWORD"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
	})
	rmq, err := repository.NewRabbitMQ(repository.RabbitMQConfig{
		Host: viper.GetString("rabbitmq.host"),
		Port: viper.GetString("rabbitmq.port"),
		Username: viper.GetString("rabbitmq.username"),
		Password: os.Getenv("RABBITMQ_PASSWORD"),
	})
	if err != nil {
		logrus.Fatalf("failed to initialize db: %s", err.Error())
	}
	repos := repository.NewRepository(db, rmq)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(store.Server)
	if err := srv.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
		logrus.Fatalf("error occured while running http server: %s", err.Error())
	}
}

func initConfig() error {
	viper.AddConfigPath("c:\\Users\\Reno\\Desktop\\My\\Projects\\Go\\UpdStore\\configs")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
