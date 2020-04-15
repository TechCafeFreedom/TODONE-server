package repository

import "github.com/gin-gonic/gin"

// MasterTxManager トランザクションマネージャ
type MasterTxManager interface {
	Transaction(ctx *gin.Context, f func(ctx *gin.Context, masterTx MasterTx) error) error
}

// MasterTx トランザクション
type MasterTx interface {
	Commit() error
	Rollback() error
}
