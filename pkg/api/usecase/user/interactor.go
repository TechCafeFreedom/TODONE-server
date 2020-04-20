package user

import (
	"context"
	"todone/pkg/domain/entity"
	"todone/pkg/domain/repository"
	userservice "todone/pkg/domain/service/user"
)

type Interactor interface {
	CreateNewUser(ctx context.Context, uid, title, description string) error
	GetUserProfile(ctx context.Context, uid string) (*entity.User, error)
	GetAll(ctx context.Context) (entity.UserSlice, error)
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

func (i *intereractor) CreateNewUser(ctx context.Context, uid, title, description string) error {
	err := i.masterTxManager.Transaction(ctx, func(ctx context.Context, masterTx repository.MasterTx) error {
		// 新規ユーザ作成
		if err := i.userService.CreateNewUser(ctx, masterTx, uid, title, description); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}

func (i *intereractor) GetUserProfile(ctx context.Context, uid string) (*entity.User, error) {
	var userData *entity.User
	var err error

	err = i.masterTxManager.Transaction(ctx, func(ctx context.Context, masterTx repository.MasterTx) error {
		// ログイン済ユーザのプロフィール情報取得
		userData, err = i.userService.GetByUID(ctx, masterTx, uid)
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

func (i *intereractor) GetAll(ctx context.Context) (entity.UserSlice, error) {
	var userSlice entity.UserSlice
	var err error

	err = i.masterTxManager.Transaction(ctx, func(ctx context.Context, masterTx repository.MasterTx) error {
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
