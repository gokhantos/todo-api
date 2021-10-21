package service

import (
	"fmt"
	"todo-api/model"
	"todo-api/repository"
)

type Service interface {
	GetTasks() interface{}
	FindTask(id string) interface{}
	CreateTask(i interface{}) interface{}
	UpdateTask(id string, i interface{}) interface{}
	DeleteTask(id string)
}

type TodoService struct {
	repository repository.Repository
}

func NewTodoService(repository repository.Repository) Service {
	return TodoService{repository}
}

func (s TodoService) GetTasks() interface{} {
	return s.repository.GetTasks()
}

func (s TodoService) FindTask(id string) interface{} {
	found := s.repository.FindTask(id)
	if found == nil {
		return "No todo found with provided id!"
	}
	return found
}

func (s TodoService) CreateTask(i interface{}) interface{} {
	request, ok := i.(model.CreateTodoRequest)
	fmt.Print("REQUEST TASK : ")
	fmt.Println(request)
	if !ok {
		return "Invalid request!"
	}

	if request.Task == "" {
		return "Task cannot be empty!"
	}

	todo := model.Todo{
		ID:     request.ID,
		Task:   request.Task,
		Status: false,
	}
	return s.repository.CreateTask(todo)
}

func (s TodoService) UpdateTask(id string, i interface{}) interface{} {
	request, ok := i.(model.UpdateTodoRequest)
	if !ok {
		return "Invalid request!"
	}
	fmt.Print(request)

	found := s.repository.FindTask(id)
	if found == nil {
		return "No todo found with provided id!"
	}
	todo := found.(model.Todo)
	return s.repository.UpdateTask(todo)
}

func (s TodoService) DeleteTask(id string) {
	s.repository.DeleteTask(id)
}
