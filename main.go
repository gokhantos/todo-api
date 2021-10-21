package main

import (
	"context"
	"fmt"
	"log"
	"todo-api/handler"
	"todo-api/repository"
	"todo-api/service"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func SetupMongoDB(mode string) *mongo.Collection {
	clientOptions := options.Client().ApplyURI("mongodb://root:pass@localhost:27017")
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB...")

	if mode == "test" {
		client.Database(mode).Drop(context.Background())
	}
	return client.Database(mode).Collection("todos")
}

func server(mode string) {
	if mode == "test" {
		todoRepository := repository.NewTodoRepository(SetupMongoDB(mode))
		todoService := service.NewTodoService(todoRepository)
		todoHandler := handler.NewTodoHandler(todoService)
		e := echo.New()
		e.Use(middleware.CORS())
		e.GET("/api/todos", todoHandler.HandleGetTasks)
		e.POST("/api/todos", todoHandler.HandleCreateTask)
		e.PUT("/api/todos/:id", todoHandler.HandleUpdateTask)
		e.DELETE("/api/todos/:id", todoHandler.HandleDeleteTask)
		e.Logger.Fatal(e.Start(":8081"))
	} else if mode == "main" {
		todoRepository := repository.NewTodoRepository(SetupMongoDB(mode))
		todoService := service.NewTodoService(todoRepository)
		todoHandler := handler.NewTodoHandler(todoService)

		e := echo.New()
		e.Use(middleware.CORS())
		e.GET("/api/todos", todoHandler.HandleGetTasks)
		e.POST("/api/todos", todoHandler.HandleCreateTask)
		e.PUT("/api/todos/:id", todoHandler.HandleUpdateTask)
		e.DELETE("/api/todos/:id", todoHandler.HandleDeleteTask)

		e.Logger.Fatal(e.Start(":8080"))
	}
}

func main() {
	server("main")
}
