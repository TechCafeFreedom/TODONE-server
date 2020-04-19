package user

import (
	"todone/db/mysql/model"
	"todone/pkg/domain/entity"
	"todone/pkg/domain/repository"
	"todone/pkg/domain/repository/user"
	"todone/pkg/infrastructure/mysql"

	"github.com/gin-gonic/gin"
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

func (u userRepositoryImpliment) InsertUser(ctx *gin.Context, masterTx repository.MasterTx, uid, name, thumbnail string) error {
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

func (u userRepositoryImpliment) SelectByPK(ctx *gin.Context, masterTx repository.MasterTx, userID int) (*entity.User, error) {
	exec, err := mysql.ExtractExecutor(masterTx)
	if err != nil {
		return nil, err
	}
	userData, err := model.FindUser(ctx, exec, userID)
	if err != nil {
		return nil, err
	}

	return entity.ConvertToUserEntity(userData), nil
}

func (u userRepositoryImpliment) SelectByUID(ctx *gin.Context, masterTx repository.MasterTx, uid string) (*entity.User, error) {
	exec, err := mysql.ExtractExecutor(masterTx)
	if err != nil {
		return nil, err
	}
	userData, err := model.Users(model.UserWhere.UID.EQ(uid)).One(ctx, exec)
	if err != nil {
		return nil, err
	}

	return entity.ConvertToUserEntity(userData), nil
}

func (u userRepositoryImpliment) SelectAll(ctx *gin.Context, masterTx repository.MasterTx) (entity.UserSlice, error) {
	exec, err := mysql.ExtractExecutor(masterTx)
	if err != nil {
		return nil, err
	}
	queries := []qm.QueryMod{}
	users, err := model.Users(queries...).All(ctx, exec)
	if err != nil {
		return nil, err
	}

	return entity.ConvertToUserSliceEntity(users), nil
}
