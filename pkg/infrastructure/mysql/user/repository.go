package user

import (
	"context"
	"todone/db/mysql/model"
	"todone/pkg/domain/entity"
	"todone/pkg/domain/repository"
	"todone/pkg/domain/repository/user"
	"todone/pkg/infrastructure/mysql"
	"todone/pkg/terrors"

	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type userRepositoryImpliment struct {
	masterTxManager repository.MasterTxManager
}

func New(masterTxManager repository.MasterTxManager) user.Repository {
	return &userRepositoryImpliment{
		masterTxManager: masterTxManager,
	}
}

func (u userRepositoryImpliment) InsertUser(ctx context.Context, masterTx repository.MasterTx, uid, name, thumbnail string) error {
	newUserData := &model.User{
		UID:       uid,
		Name:      name,
		Thumbnail: null.StringFrom(thumbnail),
	}

	exec, err := mysql.ExtractExecutor(masterTx)
	if err != nil {
		return terrors.Stack(err)
	}
	if err := newUserData.Insert(ctx, exec, boil.Infer()); err != nil {
		return terrors.Stack(err)
	}

	return nil
}

func (u userRepositoryImpliment) SelectByPK(ctx context.Context, masterTx repository.MasterTx, userID int) (*entity.User, error) {
	exec, err := mysql.ExtractExecutor(masterTx)
	if err != nil {
		return nil, terrors.Stack(err)
	}
	userData, err := model.FindUser(ctx, exec, userID)
	if err != nil {
		return nil, terrors.Stack(err)
	}

	return entity.ConvertToUserEntity(userData), nil
}

func (u userRepositoryImpliment) SelectByUID(ctx context.Context, masterTx repository.MasterTx, uid string) (*entity.User, error) {
	exec, err := mysql.ExtractExecutor(masterTx)
	if err != nil {
		return nil, terrors.Stack(err)
	}
	userData, err := model.Users(model.UserWhere.UID.EQ(uid)).One(ctx, exec)
	if err != nil {
		return nil, terrors.Stack(err)
	}

	return entity.ConvertToUserEntity(userData), nil
}

func (u userRepositoryImpliment) SelectAll(ctx context.Context, masterTx repository.MasterTx) (entity.UserSlice, error) {
	exec, err := mysql.ExtractExecutor(masterTx)
	if err != nil {
		return nil, terrors.Stack(err)
	}
	queries := []qm.QueryMod{}
	users, err := model.Users(queries...).All(ctx, exec)
	if err != nil {
		return nil, terrors.Stack(err)
	}

	return entity.ConvertToUserSliceEntity(users), nil
}
