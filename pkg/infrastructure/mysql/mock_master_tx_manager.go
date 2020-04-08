package mysql

import (
	"github.com/gin-gonic/gin"
)

// TxnManagerは構造が特殊だったのでmockを自前で実装。
type MockMasterTxManager struct {
	masterTx MasterTx
}

func NewMockMasterTxManager(masterTx MasterTx) MasterTxManager {
	return &MockMasterTxManager{masterTx}
}

func (m *MockMasterTxManager) Transaction(ctx *gin.Context, f func(ctx *gin.Context, masterTx MasterTx) error) error {
	return f(ctx, m.masterTx)
}

type MockMasterTx struct {
}

func NewMockMasterTx() MasterTx {
	return &MockMasterTx{}
}

func (m *MockMasterTx) Commit() error {
	return nil
}

func (m *MockMasterTx) Rollback() error {
	return nil
}
