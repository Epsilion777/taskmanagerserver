package main

import (
	"context"
	"os"
	"taskmanagerserver/internal/adapters/storage/postgres"
	todostorage "taskmanagerserver/internal/adapters/storage/postgres/todo_storage"
	"taskmanagerserver/internal/delivery/rest"
	"taskmanagerserver/internal/usecase/todo"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {

	ctx := context.Background()
	connString := os.Args[1]
	// connString := `postgres://postgres:123321@localhost/test?sslmode=disable`

	postgresDB := postgres.InitPostgres(ctx, connString)
	defer postgresDB.Close()

	todoStorage := todostorage.NewStorage(postgresDB)
	todoUseCase := todo.NewUsecase(todoStorage)

	router := gin.Default()

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:3000"}
	router.Use(cors.New(config))

	rest.NewHandler(todoUseCase, router)

	router.Run(":8080")
}
