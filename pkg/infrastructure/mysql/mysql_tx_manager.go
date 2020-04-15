package mysql

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"todone/pkg/domain/repository"

	"github.com/gin-gonic/gin"
	"github.com/volatiletech/sqlboiler/boil"
)

type dbMasterTxManager struct {
	db *sql.DB
}

func NewDBMasterTxManager(db *sql.DB) repository.MasterTxManager {
	return &dbMasterTxManager{db}
}

func (m *dbMasterTxManager) Transaction(ctx *gin.Context, f func(ctx *gin.Context, masterTx repository.MasterTx) error) error {
	tx, err := m.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer func() {
		// panic
		if p := recover(); p != nil {
			e := tx.Rollback()
			if e != nil {
				log.Fatal(fmt.Sprintf("failed to MySQL Rollback: %v", e))
			}
			panic(p) // re-throw panic after Rollback
		}
		// error
		if err != nil {
			e := tx.Rollback()
			if e != nil {
				log.Fatal(fmt.Sprintf("failed to MySQL Rollback: %v", e))
			}
			return
		}
		// 正常
		e := tx.Commit()
		if e != nil {
			log.Fatal(fmt.Sprintf("failed to MySQL Commit: %v", e))
		}

	}()
	err = f(ctx, &dbMasterTx{tx})
	if err != nil {
		return err
	}
	return nil
}

type dbMasterTx struct {
	tx *sql.Tx
}

func (m *dbMasterTx) Commit() error {
	return m.tx.Commit()
}

func (m *dbMasterTx) Rollback() error {
	return m.tx.Rollback()
}

func ExtractExecutor(masterTx repository.MasterTx) (boil.ContextExecutor, error) {
	return ExtractTx(masterTx)
}

func ExtractTx(masterTx repository.MasterTx) (*sql.Tx, error) {
	// キャストする
	tx, ok := masterTx.(*dbMasterTx)
	if !ok {
		return nil, errors.New("masterTx cannot cast to dbMasterTx")
	}
	return tx.tx, nil
}
