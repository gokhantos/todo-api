// Code generated by MockGen. DO NOT EDIT.
// Source: repository.go

// Package mock_repository is a generated GoMock package.
package repository

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockRepository is a mock of Repository interface.
type MockRepository struct {
	ctrl     *gomock.Controller
	recorder *MockRepositoryMockRecorder
}

// MockRepositoryMockRecorder is the mock recorder for MockRepository.
type MockRepositoryMockRecorder struct {
	mock *MockRepository
}

// NewMockRepository creates a new mock instance.
func NewMockRepository(ctrl *gomock.Controller) *MockRepository {
	mock := &MockRepository{ctrl: ctrl}
	mock.recorder = &MockRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRepository) EXPECT() *MockRepositoryMockRecorder {
	return m.recorder
}

// CreateTask mocks base method.
func (m *MockRepository) CreateTask(i interface{}) interface{} {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateTask", i)
	ret0, _ := ret[0].(interface{})
	return ret0
}

// CreateTask indicates an expected call of CreateTask.
func (mr *MockRepositoryMockRecorder) CreateTask(i interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateTask", reflect.TypeOf((*MockRepository)(nil).CreateTask), i)
}

// DeleteTask mocks base method.
func (m *MockRepository) DeleteTask(id string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "DeleteTask", id)
}

// DeleteTask indicates an expected call of DeleteTask.
func (mr *MockRepositoryMockRecorder) DeleteTask(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteTask", reflect.TypeOf((*MockRepository)(nil).DeleteTask), id)
}

// FindTask mocks base method.
func (m *MockRepository) FindTask(id string) interface{} {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindTask", id)
	ret0, _ := ret[0].(interface{})
	return ret0
}

// FindTask indicates an expected call of FindTask.
func (mr *MockRepositoryMockRecorder) FindTask(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindTask", reflect.TypeOf((*MockRepository)(nil).FindTask), id)
}

// GetTasks mocks base method.
func (m *MockRepository) GetTasks() interface{} {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTasks")
	ret0, _ := ret[0].(interface{})
	return ret0
}

// GetTasks indicates an expected call of GetTasks.
func (mr *MockRepositoryMockRecorder) GetTasks() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTasks", reflect.TypeOf((*MockRepository)(nil).GetTasks))
}

// UpdateTask mocks base method.
func (m *MockRepository) UpdateTask(i interface{}) interface{} {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateTask", i)
	ret0, _ := ret[0].(interface{})
	return ret0
}

// UpdateTask indicates an expected call of UpdateTask.
func (mr *MockRepositoryMockRecorder) UpdateTask(i interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateTask", reflect.TypeOf((*MockRepository)(nil).UpdateTask), i)
}
