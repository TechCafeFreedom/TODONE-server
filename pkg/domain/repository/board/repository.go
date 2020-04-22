package board

import (
	"context"
	"todone/pkg/domain/entity"
	"todone/pkg/domain/repository"
)

type Repository interface {
	InsertBoard(ctx context.Context, masterTx repository.MasterTx, userID int, title, description string) error
	SelectByPK(ctx context.Context, masterTx repository.MasterTx, id int) (*entity.Board, error)
	SelectByUserID(ctx context.Context, masterTx repository.MasterTx, userID int) (entity.BoardSlice, error)
	SelectAll(ctx context.Context, masterTx repository.MasterTx) (entity.BoardSlice, error)
}
