// Code generated by MockGen. DO NOT EDIT.
// Source: pkg/domain/service/board/service.go

// Package mock_board is a generated GoMock package.
package mock_board

import (
	reflect "reflect"
	entity "todone/pkg/domain/entity"
	repository "todone/pkg/domain/repository"

	gin "github.com/gin-gonic/gin"
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

// CreateNewBoard mocks base method
func (m *MockService) CreateNewBoard(ctx *gin.Context, masterTx repository.MasterTx, userID int, title, description string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateNewBoard", ctx, masterTx, userID, title, description)
	ret0, _ := ret[0].(error)
	return ret0
}

// CreateNewBoard indicates an expected call of CreateNewBoard
func (mr *MockServiceMockRecorder) CreateNewBoard(ctx, masterTx, userID, title, description interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateNewBoard", reflect.TypeOf((*MockService)(nil).CreateNewBoard), ctx, masterTx, userID, title, description)
}

// GetByPK mocks base method
func (m *MockService) GetByPK(ctx *gin.Context, masterTx repository.MasterTx, id int) (*entity.Board, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByPK", ctx, masterTx, id)
	ret0, _ := ret[0].(*entity.Board)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByPK indicates an expected call of GetByPK
func (mr *MockServiceMockRecorder) GetByPK(ctx, masterTx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByPK", reflect.TypeOf((*MockService)(nil).GetByPK), ctx, masterTx, id)
}

// GetByUserID mocks base method
func (m *MockService) GetByUserID(ctx *gin.Context, masterTx repository.MasterTx, userID int) (entity.BoardSlice, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByUserID", ctx, masterTx, userID)
	ret0, _ := ret[0].(entity.BoardSlice)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByUserID indicates an expected call of GetByUserID
func (mr *MockServiceMockRecorder) GetByUserID(ctx, masterTx, userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByUserID", reflect.TypeOf((*MockService)(nil).GetByUserID), ctx, masterTx, userID)
}

// GetAll mocks base method
func (m *MockService) GetAll(ctx *gin.Context, masterTx repository.MasterTx) (entity.BoardSlice, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAll", ctx, masterTx)
	ret0, _ := ret[0].(entity.BoardSlice)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAll indicates an expected call of GetAll
func (mr *MockServiceMockRecorder) GetAll(ctx, masterTx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAll", reflect.TypeOf((*MockService)(nil).GetAll), ctx, masterTx)
}