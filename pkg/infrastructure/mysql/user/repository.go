package user

import (
	"context"
	"todone/db/mysql/model"
	"todone/pkg/domain/repository/user"

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

func (u userRepositoryImpliment) InsertUser(user *model.User) error {
	if err := user.Insert(context.Background(), u.db, boil.Infer()); err != nil {
		return err
	}

	return nil
}

func (u userRepositoryImpliment) SelectByPK(userID string) (*model.User, error) {
	userData, err := model.FindUser(context.Background(), u.db, userID)
	if err != nil {
		return nil, err
	}

	return userData, nil
}

func (u userRepositoryImpliment) SelectAll() (model.UserSlice, error) {
	queries := []qm.QueryMod{}
	users, err := model.Users(queries...).All(context.Background(), u.db)
	if err != nil {
		return nil, err
	}

	return users, nil
}
