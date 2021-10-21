package service

import (
	"testing"
	"todo-api/model"
	"todo-api/repository"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func Test_GetAllTasks(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	var todoList []model.Todo
	todoList = append(todoList, model.Todo{
		ID:     "1",
		Task:   "Buy milk",
		Status: false,
	})

	repository := repository.NewMockRepository(controller)
	repository.EXPECT().GetTasks().Return(todoList).Times(1)
	service := TodoService{repository}

	all := service.GetTasks().([]model.Todo)
	assert.NotNil(t, all)
	assert.NotEmpty(t, all)
	assert.Equal(t, 1, len(all))
	assert.Equal(t, "1", all[0].ID)
	assert.Equal(t, "Buy milk", all[0].Task)
	assert.False(t, all[0].Status)
}

func Test_CreateTask(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	repository := repository.NewMockRepository(controller)
	repository.EXPECT().CreateTask(model.Todo{ID: "1", Task: "Buy milk"}).Return(model.Todo{
		ID:     "1",
		Task:   "Buy milk",
		Status: false,
	}).Times(1)

	service := NewTodoService(repository)
	created := service.CreateTask(model.CreateTodoRequest{ID: "1", Task: "Buy milk"})

	assert.NotNil(t, created)
	todo := created.(model.Todo)
	assert.Equal(t, "1", todo.ID)
	assert.Equal(t, "Buy milk", todo.Task)
	assert.False(t, todo.Status)
}

func Test_UpdateTask(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	beforeUpdate := model.Todo{
		ID:     "1",
		Task:   "Buy milk",
		Status: false,
	}

	afterUpdate := model.Todo{
		ID:     "1",
		Task:   "Buy milk",
		Status: true,
	}

	repository := repository.NewMockRepository(controller)
	repository.EXPECT().FindTask(gomock.Eq("1")).Return(beforeUpdate).Times(1)
	repository.EXPECT().UpdateTask(gomock.Eq(beforeUpdate)).Return(afterUpdate).Times(1)

	service := TodoService{repository}
	updated := service.UpdateTask("1", model.UpdateTodoRequest{
		Status: true,
	})

	assert.NotNil(t, updated)

	todo, ok := updated.(model.Todo)
	assert.True(t, ok)
	assert.Equal(t, "1", todo.ID)
	assert.True(t, todo.Status)
}

func Test_DeleteTask(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	repository := repository.NewMockRepository(controller)
	repository.EXPECT().DeleteTask(gomock.Eq("1")).Times(1)

	service := TodoService{repository}
	service.DeleteTask("1")
}
