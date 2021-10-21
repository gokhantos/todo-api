package handler

import (
	"fmt"
	"todo-api/model"
	"todo-api/service"
	util "todo-api/util"

	"github.com/labstack/echo"
)

type TodoHandler struct {
	service service.Service
}

type Handler interface {
	HandleGetTasks(ctx echo.Context) error
	HandleCreateTask(ctx echo.Context) error
	HandleUpdateTask(ctx echo.Context) error
	HandleDeleteTask(ctx echo.Context) error
}

func NewTodoHandler(service service.Service) Handler {
	return TodoHandler{service}
}

func (handler TodoHandler) HandleGetTasks(ctx echo.Context) error {
	return ctx.JSON(200, handler.service.GetTasks())
}

func (handler TodoHandler) HandleCreateTask(ctx echo.Context) error {
	var request model.CreateTodoRequest
	_ = ctx.Bind(&request)
	fmt.Print("CREATE REQUEST: ")
	fmt.Println(request)
	response := handler.service.CreateTask(request)
	error := extractError(response)
	if error != nil {
		return ctx.JSON(error.GetStatus(), error)
	}
	return ctx.JSON(201, response)
}

func (handler TodoHandler) HandleUpdateTask(ctx echo.Context) error {
	var request model.UpdateTodoRequest
	fmt.Println(request)
	_ = ctx.Bind(&request)
	id := ctx.Param("id")
	response := handler.service.UpdateTask(id, request)
	error := extractError(response)
	if error != nil {
		return ctx.JSON(error.GetStatus(), error)
	}
	return ctx.JSON(202, response)
}

func (handler TodoHandler) HandleDeleteTask(ctx echo.Context) error {
	id := ctx.Param("id")
	handler.service.DeleteTask(id)
	return ctx.JSON(200, id)
}

func extractError(i interface{}) util.Error {
	error, ok := i.(util.Error)
	if ok {
		return error
	}
	return nil
}
