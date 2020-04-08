package user

import (
	"todone/db/mysql/model"
	userservice "todone/pkg/domain/service/user"
	"todone/pkg/infrastructure/mysql"

	"github.com/gin-gonic/gin"
)

type Interactor interface {
	CreateNewUser(ctx *gin.Context, userID, title, description string) error
	GetByPK(userID string) (*model.User, error)
	GetAll() (model.UserSlice, error)
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
		if err := i.userService.CreateNewUser(userID, title, description); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}

func (i *intereractor) GetByPK(userID string) (*model.User, error) {
	user, err := i.userService.GetByPK(userID)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (i *intereractor) GetAll() (model.UserSlice, error) {
	users, err := i.userService.GetAll()
	if err != nil {
		return nil, err
	}
	return users, nil
}
