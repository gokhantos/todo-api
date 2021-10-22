package controller

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"todo-api/model"
	"todo-api/service"

	"github.com/golang/mock/gomock"
	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
)

func Test_GetTodos(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	var todos []model.Todo
	todos = append(todos, model.Todo{
		ID:     "1",
		Task:   "Buy milk",
		Status: false,
	})

	service := service.NewMockService(controller)
	service.EXPECT().GetTasks().Return(todos).Times(1)
	handler := TodoController{service}

	expected := `[{"id":"1","task":"Buy milk","status":false}]
`

	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/todos", nil)
	req.Header.Set(echo.HeaderAccept, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	ctx := e.NewContext(req, rec)

	if assert.NoError(t, handler.ControlGetTasks(ctx)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, expected, rec.Body.String())
	}
}

func Test_ShouldCreateTodo(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	request := model.CreateTodoRequest{Task: "Buy milk"}
	requestAsString := `{"task":"Buy milk"}`
	expected := model.Todo{
		ID:     "1",
		Task:   "Buy milk",
		Status: false,
	}
	expectedAsString := `{"id":"1","task":"Buy milk","status":false}
`

	service := service.NewMockService(controller)
	service.EXPECT().CreateTask(gomock.Eq(request)).Return(expected).Times(1)

	handler := TodoController{service}
	e := echo.New()
	req := httptest.NewRequest(http.MethodPost, "/todos", strings.NewReader(requestAsString))
	req.Header.Set(echo.HeaderAccept, echo.MIMEApplicationJSON)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	ctx := e.NewContext(req, rec)

	if assert.NoError(t, handler.ControlCreateTask(ctx)) {
		assert.Equal(t, http.StatusCreated, rec.Code)
		assert.Equal(t, expectedAsString, rec.Body.String())
	}
}

func Test_UpdateTask(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	request := model.UpdateTodoRequest{
		Status: true,
	}

	requestAsString := `{"task":"Buy milk","status":true}
`
	expected := model.Todo{
		ID:     "1",
		Task:   "Buy milk",
		Status: true,
	}
	expectedAsString := `{"id":"1","task":"Buy milk","status":true}
`
	service := service.NewMockService(controller)
	service.EXPECT().UpdateTask(gomock.Eq("1"), gomock.Eq(request)).Return(expected).Times(1)
	handler := TodoController{service}

	e := echo.New()
	req := httptest.NewRequest(http.MethodPut, "/todos/1", strings.NewReader(requestAsString))
	req.Header.Set(echo.HeaderAccept, echo.MIMEApplicationJSON)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	ctx := e.NewContext(req, rec)
	ctx.SetParamNames("id")
	ctx.SetParamValues("1")

	if assert.NoError(t, handler.ControlUpdateTask(ctx)) {
		assert.Equal(t, http.StatusAccepted, rec.Code)
		assert.Equal(t, expectedAsString, rec.Body.String())
	}
}

func Test_DeleteTask(t *testing.T) {
	controller := gomock.NewController(t)
	defer controller.Finish()

	service := service.NewMockService(controller)
	service.EXPECT().DeleteTask(gomock.Eq("1"))
	handler := TodoController{service}

	e := echo.New()
	req := httptest.NewRequest(http.MethodDelete, "/todos/1", nil)
	rec := httptest.NewRecorder()
	ctx := e.NewContext(req, rec)
	ctx.SetParamNames("id")
	ctx.SetParamValues("1")

	if assert.NoError(t, handler.ControlDeleteTask(ctx)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}
}
