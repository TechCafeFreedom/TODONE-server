package user

import (
	"todone/db/mysql/model"
	userservice "todone/pkg/domain/service/user"
	"todone/pkg/infrastructure/mysql"

	"github.com/gin-gonic/gin"
)

type Interactor interface {
	CreateNewUser(ctx *gin.Context, userID, title, description string) error
	GetByPK(ctx *gin.Context, userID string) (*model.User, error)
	GetAll(ctx *gin.Context) (model.UserSlice, error)
}

type intereractor struct {
	masterTxManager mysql.MasterTxManager
	userService     userservice.Service
}

func New(masterTxManager mysql.MasterTxManager, userService userservice.Service) Interactor {
	return &intereractor{
		masterTxManager: masterTxManager,
		userService:     userService,
	}
}

func (i *intereractor) CreateNewUser(ctx *gin.Context, userID, title, description string) error {
	err := i.masterTxManager.Transaction(ctx, func(ctx *gin.Context, masterTx mysql.MasterTx) error {
		if err := i.userService.CreateNewUser(ctx, masterTx, userID, title, description); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}

func (i *intereractor) GetByPK(ctx *gin.Context, userID string) (*model.User, error) {
	var userData *model.User
	var err error

	err = i.masterTxManager.Transaction(ctx, func(ctx *gin.Context, masterTx mysql.MasterTx) error {
		userData, err = i.userService.GetByPK(ctx, masterTx, userID)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return userData, nil
}

func (i *intereractor) GetAll(ctx *gin.Context) (model.UserSlice, error) {
	var userSlice model.UserSlice
	var err error

	err = i.masterTxManager.Transaction(ctx, func(ctx *gin.Context, masterTx mysql.MasterTx) error {
		userSlice, err = i.userService.GetAll(ctx, masterTx)
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return userSlice, nil
}
