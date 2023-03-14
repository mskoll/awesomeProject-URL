package main

import (
	"awesomeProject-URL/internal"
	"awesomeProject-URL/internal/handler"
	"awesomeProject-URL/internal/repository"
	"awesomeProject-URL/internal/service"
	"context"
	"flag"
	"github.com/spf13/viper"
	"log"
	"os"
	"os/signal"
	"time"
)

func init() {
	flagDB := flag.Bool("d", false, "data from the database")
	flag.Parse()
	viper.Set("flagDB", flagDB)
}

func main() {

	// инициализация конфига
	if err := initConfig(); err != nil {
		log.Fatalf("Config error: %s", err.Error())
	}

	// подключение к БД
	// передаем данные для подключения
	db, err := repository.InitDB(repository.Conf{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		User:     viper.GetString("db.user"),
		Password: viper.GetString("db.password"),
		DBName:   viper.GetString("db.dbname"),
	})
	if err != nil {
		log.Fatalf("DB-init error: %s", err.Error())
	}
	log.Printf("DB connected\n")

	// инициализация репозитория для работы с БД
	repositories := repository.NewRepository(db)
	// инициализация сервиса - бизнес-логика
	services := service.NewService(repositories)
	// инициализация хэндлера
	handlers := handler.NewHandler(services)

	// инициализация сервера
	server := new(internal.Server)

	// запуск сервера
	go func() {
		if err := server.Run(viper.GetString("port"), handlers.InitRoutes()); err != nil {
			log.Printf("Server error: %s", err.Error())
		}
	}()
	log.Printf("Server started\n")
	// канал для получения сигналов системы
	stop := make(chan os.Signal, 1)
	// получение сигнала, что приложение завершилось
	signal.Notify(stop, os.Interrupt)
	<-stop

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Printf("Server shutdown error: %s", err.Error())
	}

	if err := db.Close(); err != nil {
		log.Printf("DB connection close error: %s", err.Error())
	}
}

// initConfig иницаиализация конфига
func initConfig() error {
	//viper.AddConfigPath("internal/config") for IDE
	viper.AddConfigPath("../internal/config")
	viper.SetConfigName("config")
	return viper.ReadInConfig()
}
