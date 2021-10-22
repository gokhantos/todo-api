package main

import (
	"context"
	"fmt"
	"log"
	"todo-api/controller"
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
		TodoController := controller.NewTodoController(todoService)
		e := echo.New()
		e.Use(middleware.CORS())
		e.GET("/api/todos", TodoController.ControlGetTasks)
		e.POST("/api/todos", TodoController.ControlCreateTask)
		e.PUT("/api/todos/:id", TodoController.ControlUpdateTask)
		e.DELETE("/api/todos/:id", TodoController.ControlDeleteTask)
		e.Logger.Fatal(e.Start(":8081"))
	} else if mode == "main" {
		todoRepository := repository.NewTodoRepository(SetupMongoDB(mode))
		todoService := service.NewTodoService(todoRepository)
		TodoController := controller.NewTodoController(todoService)

		e := echo.New()
		e.Use(middleware.CORS())
		e.GET("/api/todos", TodoController.ControlGetTasks)
		e.POST("/api/todos", TodoController.ControlCreateTask)
		e.PUT("/api/todos/:id", TodoController.ControlUpdateTask)
		e.DELETE("/api/todos/:id", TodoController.ControlDeleteTask)

		e.Logger.Fatal(e.Start(":8080"))
	}
}

func main() {
	server("main")
}
