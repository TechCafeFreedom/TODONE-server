// Code generated by MockGen. DO NOT EDIT.
// Source: pkg/domain/service/user/service.go

// Package mock_user is a generated GoMock package.
package mock_user

import (
	reflect "reflect"
	models "todone/db/mysql/models"

	gomock "github.com/golang/mock/gomock"
)

// MockService is a mock of Service interface
type MockService struct {
	ctrl     *gomock.Controller
	recorder *MockServiceMockRecorder
}

// MockServiceMockRecorder is the mock recorder for MockService
type MockServiceMockRecorder struct {
	mock *MockService
}

// NewMockService creates a new mock instance
func NewMockService(ctrl *gomock.Controller) *MockService {
	mock := &MockService{ctrl: ctrl}
	mock.recorder = &MockServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockService) EXPECT() *MockServiceMockRecorder {
	return m.recorder
}

// Login mocks base method
func (m *MockService) Login(email, password string) (*models.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Login", email, password)
	ret0, _ := ret[0].(*models.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Login indicates an expected call of Login
func (mr *MockServiceMockRecorder) Login(email, password interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Login", reflect.TypeOf((*MockService)(nil).Login), email, password)
}

// GetUser mocks base method
func (m *MockService) GetUser(id string) ([]*models.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUser", id)
	ret0, _ := ret[0].([]*models.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUser indicates an expected call of GetUser
func (mr *MockServiceMockRecorder) GetUser(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUser", reflect.TypeOf((*MockService)(nil).GetUser), id)
}

// CreateNewUser mocks base method
func (m *MockService) CreateNewUser(newUser *models.User) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateNewUser", newUser)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateNewUser indicates an expected call of CreateNewUser
func (mr *MockServiceMockRecorder) CreateNewUser(newUser interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateNewUser", reflect.TypeOf((*MockService)(nil).CreateNewUser), newUser)
}

// SelectAll mocks base method
func (m *MockService) SelectAll() ([]*models.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SelectAll")
	ret0, _ := ret[0].([]*models.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SelectAll indicates an expected call of SelectAll
func (mr *MockServiceMockRecorder) SelectAll() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SelectAll", reflect.TypeOf((*MockService)(nil).SelectAll))
}