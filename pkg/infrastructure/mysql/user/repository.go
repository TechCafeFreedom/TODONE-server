package user

import (
	"todone/db/mysql/model"
	"todone/pkg/domain/repository"
	"todone/pkg/domain/repository/user"
	"todone/pkg/infrastructure/mysql"

	"github.com/gin-gonic/gin"
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries/qm"
)

type userRepositoryImpliment struct {
	db boil.ContextExecutor
}

func New(db boil.ContextExecutor) user.Repository {
	return &userRepositoryImpliment{
		db: db,
	}
}

func (u userRepositoryImpliment) InsertUser(ctx *gin.Context, masterTx repository.MasterTx, userData *model.User) error {
	exec, err := mysql.ExtractExecutor(masterTx)
	if err != nil {
		return err
	}
	if err := userData.Insert(ctx, exec, boil.Infer()); err != nil {
		return err
	}

	return nil
}

func (u userRepositoryImpliment) SelectByPK(ctx *gin.Context, masterTx repository.MasterTx, userID string) (*model.User, error) {
	exec, err := mysql.ExtractExecutor(masterTx)
	if err != nil {
		return nil, err
	}
	userData, err := model.FindUser(ctx, exec, userID)
	if err != nil {
		return nil, err
	}

	return userData, nil
}

func (u userRepositoryImpliment) SelectAll(ctx *gin.Context, masterTx repository.MasterTx) (model.UserSlice, error) {
	exec, err := mysql.ExtractExecutor(masterTx)
	if err != nil {
		return nil, err
	}
	queries := []qm.QueryMod{}
	users, err := model.Users(queries...).All(ctx, exec)
	if err != nil {
		return nil, err
	}

	return users, nil
}
