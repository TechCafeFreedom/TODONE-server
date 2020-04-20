package repository

import "context"

// MasterTxManager トランザクションマネージャ
type MasterTxManager interface {
	Transaction(ctx context.Context, f func(ctx context.Context, masterTx MasterTx) error) error
}

// MasterTx トランザクション
type MasterTx interface {
	Commit() error
	Rollback() error
}
