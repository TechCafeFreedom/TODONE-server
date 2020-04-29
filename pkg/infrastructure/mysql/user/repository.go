package user

import (
	"context"
	"todone/db/mysql/model"
	"todone/pkg/domain/entity"
	"todone/pkg/domain/repository"
	"todone/pkg/domain/repository/user"
	"todone/pkg/infrastructure/mysql"

	"github.com/volatiletech/null"
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries/qm"
)

type userRepositoryImpliment struct {
	masterTxManager repository.MasterTxManager
}

func New(masterTxManager repository.MasterTxManager) user.Repository {
	return &userRepositoryImpliment{
		masterTxManager: masterTxManager,
	}
}

func (u *userRepositoryImpliment) InsertUser(ctx context.Context, masterTx repository.MasterTx, uid, name, thumbnail string) error {
	newUserData := &model.User{
		UID:       uid,
		Name:      name,
		Thumbnail: null.StringFrom(thumbnail),
	}

	exec, err := mysql.ExtractExecutor(masterTx)
	if err != nil {
		return err
	}
	if err := newUserData.Insert(ctx, exec, boil.Infer()); err != nil {
		return err
	}

	return nil
}

func (u *userRepositoryImpliment) SelectByPK(ctx context.Context, masterTx repository.MasterTx, userID int) (*entity.User, error) {
	exec, err := mysql.ExtractExecutor(masterTx)
	if err != nil {
		return nil, err
	}
	userData, err := model.FindUser(ctx, exec, userID)
	if err != nil {
		return nil, err
	}

	return ConvertToUserEntity(userData), nil
}

func (u *userRepositoryImpliment) SelectByUID(ctx context.Context, masterTx repository.MasterTx, uid string) (*entity.User, error) {
	exec, err := mysql.ExtractExecutor(masterTx)
	if err != nil {
		return nil, err
	}
	userData, err := model.Users(model.UserWhere.UID.EQ(uid)).One(ctx, exec)
	if err != nil {
		return nil, err
	}

	return ConvertToUserEntity(userData), nil
}

func (u *userRepositoryImpliment) SelectAll(ctx context.Context, masterTx repository.MasterTx) (entity.UserSlice, error) {
	exec, err := mysql.ExtractExecutor(masterTx)
	if err != nil {
		return nil, err
	}
	queries := []qm.QueryMod{}
	users, err := model.Users(queries...).All(ctx, exec)
	if err != nil {
		return nil, err
	}

	return ConvertToUserSliceEntity(users), nil
}

func ConvertToUserEntity(userData *model.User) *entity.User {
	return &entity.User{
		ID:        userData.ID,
		Name:      userData.Name,
		Thumbnail: userData.Thumbnail.String,
	}
}

func ConvertToUserSliceEntity(userSlice model.UserSlice) entity.UserSlice {
	res := make(entity.UserSlice, 0, len(userSlice))
	for _, userData := range userSlice {
		res = append(res, ConvertToUserEntity(userData))
	}
	return res
}
