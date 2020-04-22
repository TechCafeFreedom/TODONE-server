package repository

import "context"

// TxnManagerは構造が特殊だったのでmockを自前で実装。
type MockMasterTxManager struct {
	masterTx MasterTx
}

func NewMockMasterTxManager(masterTx MasterTx) MasterTxManager {
	return &MockMasterTxManager{masterTx}
}

func (m *MockMasterTxManager) Transaction(ctx context.Context, f func(ctx context.Context, masterTx MasterTx) error) error {
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
