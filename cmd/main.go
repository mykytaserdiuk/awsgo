package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"github.com/mykytaserdiuk/aws-go/internal/database/postgres"
	"github.com/mykytaserdiuk/aws-go/internal/handler"
	"github.com/mykytaserdiuk/aws-go/internal/repository"
	"github.com/mykytaserdiuk/aws-go/internal/service"
)

func main() {
	// os.Setenv("PORT", ":1232")
	// os.Setenv("DB_URL", "postgres://postgres:postgres@localhost:5432/postgresDB?sslmode=disable")

	context, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	dbPool, err := postgres.NewPool(context, os.Getenv("DB_URL"))
	if err != nil {
		time.Sleep(1 * time.Second)
		log.Fatalf("Error opening db: %s", err.Error())
	}
	repo := repository.NewRepo()
	service := service.New(dbPool, repo)
	router := mux.NewRouter()
	_ = handler.NewHandler(router, service)

	server := http.Server{
		Addr:    os.Getenv("PORT"),
		Handler: router,
	}

	log.Println(os.Getenv("PORT"))
	log.Fatalln(server.ListenAndServe())
}
