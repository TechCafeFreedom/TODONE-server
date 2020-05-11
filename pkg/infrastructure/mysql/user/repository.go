package user

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"
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
	if err == sql.ErrNoRows {
		messageJP := fmt.Sprintf("ユーザが見つかりませんでした。ユーザ登録されているか確認してください。")
		messageEN := fmt.Sprintf("User not found. Please make sure signup.")
		return nil, terrors.Newf(http.StatusInternalServerError, messageJP, messageEN)
	}
	if err != nil {
		return nil, terrors.Wrapf(err, http.StatusInternalServerError, "DBアクセス時にエラーが発生しました。", "Error occured in DB access.")
	}

	return entity.ConvertToUserEntity(userData), nil
}

func (u userRepositoryImpliment) SelectByUID(ctx context.Context, masterTx repository.MasterTx, uid string) (*entity.User, error) {
	exec, err := mysql.ExtractExecutor(masterTx)
	if err != nil {
		return nil, terrors.Stack(err)
	}
	userData, err := model.Users(model.UserWhere.UID.EQ(uid)).One(ctx, exec)
	if err == sql.ErrNoRows {
		messageJP := fmt.Sprintf("不正なユーザです。")
		messageEN := fmt.Sprintf("Invalid user.")
		return nil, terrors.Newf(http.StatusInternalServerError, messageJP, messageEN)
	}
	if err != nil {
		return nil, terrors.Wrapf(err, http.StatusInternalServerError, "DBアクセス時にエラーが発生しました。", "Error occured in DB access.")
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
	if err == sql.ErrNoRows {
		messageJP := fmt.Sprintf("ユーザは1人も登録されていません。")
		messageEN := fmt.Sprintf("User doesn't exists.")
		return nil, terrors.Newf(http.StatusInternalServerError, messageJP, messageEN)
	}
	if err != nil {
		return nil, terrors.Wrapf(err, http.StatusInternalServerError, "DBアクセス時にエラーが発生しました。", "Error occured in DB access.")
	}

	return entity.ConvertToUserSliceEntity(users), nil
}
