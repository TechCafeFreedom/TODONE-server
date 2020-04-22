package user

import (
	"context"
	"todone/pkg/domain/entity"
	"todone/pkg/domain/repository"
)

type Repository interface {
	InsertUser(ctx context.Context, masterTx repository.MasterTx, uid, name, thumbnail string) error
	SelectByPK(ctx context.Context, masterTx repository.MasterTx, userID int) (*entity.User, error)
	SelectByUID(ctx context.Context, masterTx repository.MasterTx, uid string) (*entity.User, error)
	SelectAll(ctx context.Context, masterTx repository.MasterTx) (entity.UserSlice, error)
}
