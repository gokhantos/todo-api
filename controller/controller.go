package controller

import (
	"fmt"
	"todo-api/model"
	"todo-api/service"
	util "todo-api/util"

	"github.com/labstack/echo"
)

type TodoController struct {
	service service.Service
}

type Controller interface {
	ControlGetTasks(ctx echo.Context) error
	ControlCreateTask(ctx echo.Context) error
	ControlUpdateTask(ctx echo.Context) error
	ControlDeleteTask(ctx echo.Context) error
}

func NewTodoController(service service.Service) Controller {
	return TodoController{service}
}

func (controller TodoController) ControlGetTasks(ctx echo.Context) error {
	return ctx.JSON(200, controller.service.GetTasks())
}

func (controller TodoController) ControlCreateTask(ctx echo.Context) error {
	var request model.CreateTodoRequest
	_ = ctx.Bind(&request)
	fmt.Print("CREATE REQUEST: ")
	fmt.Println(request)
	response := controller.service.CreateTask(request)
	error := extractError(response)
	if error != nil {
		return ctx.JSON(error.GetStatus(), error)
	}
	return ctx.JSON(201, response)
}

func (controller TodoController) ControlUpdateTask(ctx echo.Context) error {
	var request model.UpdateTodoRequest
	fmt.Println(request)
	_ = ctx.Bind(&request)
	id := ctx.Param("id")
	response := controller.service.UpdateTask(id, request)
	error := extractError(response)
	if error != nil {
		return ctx.JSON(error.GetStatus(), error)
	}
	return ctx.JSON(202, response)
}

func (controller TodoController) ControlDeleteTask(ctx echo.Context) error {
	id := ctx.Param("id")
	controller.service.DeleteTask(id)
	return ctx.JSON(200, id)
}

func extractError(i interface{}) util.Error {
	error, ok := i.(util.Error)
	if ok {
		return error
	}
	return nil
}
