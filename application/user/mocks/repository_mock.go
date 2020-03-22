// Code generated by MockGen. DO NOT EDIT.
// Source: ./application/user/repository.go

// Package mock_user is a generated GoMock package.
package mocks

import (
	models "github.com/go-park-mail-ru/2020_1_k-on/application/models"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockRepository is a mock of Repository interface
type MockRepository struct {
	ctrl     *gomock.Controller
	recorder *MockRepositoryMockRecorder
}

// MockRepositoryMockRecorder is the mock recorder for MockRepository
type MockRepositoryMockRecorder struct {
	mock *MockRepository
}

// NewMockRepository creates a new mock instance
func NewMockRepository(ctrl *gomock.Controller) *MockRepository {
	mock := &MockRepository{ctrl: ctrl}
	mock.recorder = &MockRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockRepository) EXPECT() *MockRepositoryMockRecorder {
	return m.recorder
}

// Add mocks base method
func (m *MockRepository) Add(user *models.User) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Add", user)
	ret0, _ := ret[0].(error)
	return ret0
}

// Add indicates an expected call of Add
func (mr *MockRepositoryMockRecorder) Add(user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Add", reflect.TypeOf((*MockRepository)(nil).Add), user)
}

// Update mocks base method
func (m *MockRepository) Update(id int64, upUser *models.User) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", id, upUser)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update
func (mr *MockRepositoryMockRecorder) Update(id, upUser interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockRepository)(nil).Update), id, upUser)
}

// GetById mocks base method
func (m *MockRepository) GetById(id int64) (*models.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetById", id)
	ret0, _ := ret[0].(*models.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetById indicates an expected call of GetById
func (mr *MockRepositoryMockRecorder) GetById(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetById", reflect.TypeOf((*MockRepository)(nil).GetById), id)
}

// GetByName mocks base method
func (m *MockRepository) GetByName(login string) (*models.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByName", login)
	ret0, _ := ret[0].(*models.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByName indicates an expected call of GetByName
func (mr *MockRepositoryMockRecorder) GetByName(login interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByName", reflect.TypeOf((*MockRepository)(nil).GetByName), login)
}

// Contains mocks base method
func (m *MockRepository) Contains(login string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Contains", login)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Contains indicates an expected call of Contains
func (mr *MockRepositoryMockRecorder) Contains(login interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Contains", reflect.TypeOf((*MockRepository)(nil).Contains), login)
}