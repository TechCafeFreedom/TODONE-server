package board

import (
	"testing"
	"todone/db/mysql/model"
	"todone/pkg/domain/entity"
	"todone/pkg/domain/repository"
	"todone/pkg/domain/service/board/mock_board"
	"todone/pkg/domain/service/user/mock_user"

	"github.com/volatiletech/null"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

const (
	id          = 1
	userID      = 1
	userName    = "鈴木"
	thumbnail   = "http://s3.com/hoge.png"
	accessToken = "accessToken"
	title       = "title"
	description = "description"
)

var (
	authedUser = &model.User{
		ID:          userID,
		AccessToken: accessToken,
		Name:        "name",
		Thumbnail:   null.String{},
	}
)

func TestIntereractor_CreateNewBoard(t *testing.T) {
	ctx := &gin.Context{}
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	masterTx := repository.NewMockMasterTx()
	masterTxManager := repository.NewMockMasterTxManager(masterTx)

	boardService := mock_board.NewMockService(ctrl)
	boardService.EXPECT().CreateNewBoard(ctx, masterTx, userID, title, description).Return(nil).Times(1)

	userService := mock_user.NewMockService(ctrl)
	userService.EXPECT().GetByAccessToken(ctx, masterTx, accessToken).Return(authedUser, nil).Times(1)

	interactor := New(masterTxManager, boardService, userService)
	err := interactor.CreateNewBoard(ctx, accessToken, title, description)

	assert.NoError(t, err)
}

func TestIntereractor_GetByPK(t *testing.T) {
	ctx := &gin.Context{}
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	existedBoard := &entity.Board{
		ID:          id,
		Title:       title,
		Description: description,
		Author: &entity.User{
			ID:        userID,
			Name:      userName,
			Thumbnail: thumbnail,
		},
	}

	masterTx := repository.NewMockMasterTx()
	masterTxManager := repository.NewMockMasterTxManager(masterTx)

	boardService := mock_board.NewMockService(ctrl)
	boardService.EXPECT().GetByPK(ctx, masterTx, id).Return(existedBoard, nil).Times(1)

	userService := mock_user.NewMockService(ctrl)

	interactor := New(masterTxManager, boardService, userService)
	boards, err := interactor.GetBoardDetail(ctx, id)

	assert.NoError(t, err)
	assert.NotNil(t, boards)
}

func TestIntereractor_GetByUserID(t *testing.T) {
	ctx := &gin.Context{}
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	userBoards := entity.BoardSlice{
		{
			ID:          id,
			Title:       title,
			Description: description,
			Author: &entity.User{
				ID:        userID,
				Name:      userName,
				Thumbnail: thumbnail,
			},
		},
	}

	masterTx := repository.NewMockMasterTx()
	masterTxManager := repository.NewMockMasterTxManager(masterTx)

	boardService := mock_board.NewMockService(ctrl)
	boardService.EXPECT().GetByUserID(ctx, masterTx, userID).Return(userBoards, nil).Times(1)

	userService := mock_user.NewMockService(ctrl)
	userService.EXPECT().GetByAccessToken(ctx, masterTx, accessToken).Return(authedUser, nil).Times(1)

	interactor := New(masterTxManager, boardService, userService)
	boards, err := interactor.GetUserBoards(ctx, accessToken)

	assert.NoError(t, err)
	assert.NotNil(t, boards)
}

func TestIntereractor_GetAll(t *testing.T) {
	ctx := &gin.Context{}
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	existedBoards := entity.BoardSlice{
		{
			ID:          id,
			Title:       title,
			Description: description,
			Author: &entity.User{
				ID:        userID,
				Name:      userName,
				Thumbnail: thumbnail,
			},
		},
	}

	masterTx := repository.NewMockMasterTx()
	masterTxManager := repository.NewMockMasterTxManager(masterTx)

	boardService := mock_board.NewMockService(ctrl)
	boardService.EXPECT().GetAll(ctx, masterTx).Return(existedBoards, nil).Times(1)

	userService := mock_user.NewMockService(ctrl)

	interactor := New(masterTxManager, boardService, userService)
	boards, err := interactor.GetAll(ctx)

	assert.NoError(t, err)
	assert.NotNil(t, boards)
}
