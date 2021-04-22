package main

import (
	"backend_consumer/pkg/handler"
	"backend_consumer/pkg/repository"
	"backend_consumer/pkg/service"
	"backend_consumer/server"
	"context"
	_ "github.com/lib/pq"
	"log"
	"os"
	"os/signal"
	"syscall"
)

// @title Товарный рынок API
// @version 1.0
// description API Server для Веб-приложения "Товарный рынок"

// @host localhost:8081
// @BasePath /api

func main() {


	db, err := repository.NewPostgresDb(repository.Config{
		Host: "localhost",
		Port: "5432",
		Username: "postgres",
		Password: "password",
		DBName: "postgres",
		SSLMode: "disable",
	})

	if err != nil{
		log.Fatalf("failed to connect db: %s", err.Error())
	}

	repos := repository.NewRepository(db)
	services:= service.NewService(repos)
	handlers := handler.NewHandler(services)

	srv := new(server.Server)
	go func() {
		if err := srv.Run("8081", handlers.InitRoute()); err != nil {
			log.Fatalf("error occurred with http server: %s", err.Error())
		}
	}()

	log.Print("Backend is Started")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<- quit

	log.Print("TodoApp Shutting Down")

	if err := srv.Shutdown(context.Background()); err != nil {
		log.Printf("error occured on server shutting down: %s", err.Error())
	}

	if err := db.Close(); err != nil {
		log.Printf("error occured on db connection close: %s", err.Error())
	}
	//if err:= srv.Run("8080", handlers)
}

//func initConfig() error{
//	return ""
//}
