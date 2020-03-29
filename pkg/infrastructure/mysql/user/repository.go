package user

import (
	"context"
	model "todone/db/mysql/models"
	"todone/pkg/domain/repository/user"

	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries/qm"
	"golang.org/x/crypto/bcrypt"
)

type userRepositoryImpliment struct {
	db boil.ContextExecutor
}

func New(db boil.ContextExecutor) user.Repository {
	return &userRepositoryImpliment{
		db: db,
	}
}

func (p *userRepositoryImpliment) Login(email, password string) (*model.User, error) {
	users, err := p.SelectByEmail(email)
	if err != nil {
		return nil, err
	}

	// ユーザ認証機能
	err = bcrypt.CompareHashAndPassword([]byte(users[0].Password), []byte(password))
	if err != nil {
		return nil, err
	}

	return users[0], nil
}

func (p *userRepositoryImpliment) SelectByUserID(id string) ([]*model.User, error) {
	// ユーザ取得機能
	queries := []qm.QueryMod{
		qm.Where(model.UserColumns.UserID+"=?", id),
	}
	users, err := model.Users(queries...).All(context.Background(), p.db)
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (p *userRepositoryImpliment) InsertUser(user *model.User) error {
	// ユーザ作成機能
	if err := user.Insert(context.Background(), p.db, boil.Infer()); err != nil {
		return err
	}

	return nil
}

func (p *userRepositoryImpliment) SelectAll() ([]*model.User, error) {
	queries := []qm.QueryMod{}
	users, err := model.Users(queries...).All(context.Background(), p.db)
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (p *userRepositoryImpliment) SelectByEmail(email string) ([]*model.User, error) {
	queries := []qm.QueryMod{
		qm.Where(model.UserColumns.Email+"=?", email),
	}
	users, err := model.Users(queries...).All(context.Background(), p.db)
	if err != nil {
		return nil, err
	}
	return users, nil
}
