package user

import (
	"todone/db/mysql/model"
	"todone/pkg/domain/repository"
	userservice "todone/pkg/domain/service/user"

	"github.com/gin-gonic/gin"
)

type Interactor interface {
	CreateNewUser(ctx *gin.Context, accessToken, title, description string) error
	GetUserProfile(ctx *gin.Context, accessToken string) (*model.User, error)
	GetAll(ctx *gin.Context) (model.UserSlice, error)
}

type intereractor struct {
	masterTxManager repository.MasterTxManager
	userService     userservice.Service
}

func New(masterTxManager repository.MasterTxManager, userService userservice.Service) Interactor {
	return &intereractor{
		masterTxManager: masterTxManager,
		userService:     userService,
	}
}

func (i *intereractor) CreateNewUser(ctx *gin.Context, accessToken, title, description string) error {
	err := i.masterTxManager.Transaction(ctx, func(ctx *gin.Context, masterTx repository.MasterTx) error {
		// 新規ユーザ作成
		if err := i.userService.CreateNewUser(ctx, masterTx, accessToken, title, description); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}

func (i *intereractor) GetUserProfile(ctx *gin.Context, accessToken string) (*model.User, error) {
	var userData *model.User
	var err error

	err = i.masterTxManager.Transaction(ctx, func(ctx *gin.Context, masterTx repository.MasterTx) error {
		// ログイン済ユーザのプロフィール情報取得
		userData, err = i.userService.GetByAccessToken(ctx, masterTx, accessToken)
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

	err = i.masterTxManager.Transaction(ctx, func(ctx *gin.Context, masterTx repository.MasterTx) error {
		// (管理者用)ユーザ全件取得
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
