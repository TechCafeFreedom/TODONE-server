// Code generated by MockGen. DO NOT EDIT.
// Source: pkg/api/usecase/user/interactor.go

// Package mock_user is a generated GoMock package.
package mock_user

import (
	reflect "reflect"
	model "todone/db/mysql/model"

	gin "github.com/gin-gonic/gin"
	gomock "github.com/golang/mock/gomock"
)

// MockInteractor is a mock of Interactor interface
type MockInteractor struct {
	ctrl     *gomock.Controller
	recorder *MockInteractorMockRecorder
}

// MockInteractorMockRecorder is the mock recorder for MockInteractor
type MockInteractorMockRecorder struct {
	mock *MockInteractor
}

// NewMockInteractor creates a new mock instance
func NewMockInteractor(ctrl *gomock.Controller) *MockInteractor {
	mock := &MockInteractor{ctrl: ctrl}
	mock.recorder = &MockInteractorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockInteractor) EXPECT() *MockInteractorMockRecorder {
	return m.recorder
}

// CreateNewUser mocks base method
func (m *MockInteractor) CreateNewUser(ctx *gin.Context, userID, title, description string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateNewUser", ctx, userID, title, description)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateNewUser indicates an expected call of CreateNewUser
func (mr *MockInteractorMockRecorder) CreateNewUser(ctx, userID, title, description interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateNewUser", reflect.TypeOf((*MockInteractor)(nil).CreateNewUser), ctx, userID, title, description)
}

// GetByPK mocks base method
func (m *MockInteractor) GetByPK(userID string) (*model.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByPK", userID)
	ret0, _ := ret[0].(*model.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByPK indicates an expected call of GetByPK
func (mr *MockInteractorMockRecorder) GetByPK(userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByPK", reflect.TypeOf((*MockInteractor)(nil).GetByPK), userID)
}

// GetAll mocks base method
func (m *MockInteractor) GetAll() (model.UserSlice, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAll")
	ret0, _ := ret[0].(model.UserSlice)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAll indicates an expected call of GetAll
func (mr *MockInteractorMockRecorder) GetAll() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAll", reflect.TypeOf((*MockInteractor)(nil).GetAll))
}
