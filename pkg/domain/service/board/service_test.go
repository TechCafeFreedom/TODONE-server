package board

import (
	"testing"
	"todone/db/mysql/model"
	"todone/pkg/domain/repository"
	"todone/pkg/domain/repository/board/mock_board"

	"github.com/volatiletech/null"

	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

const (
	id          = 1
	userID      = 1
	title       = "title"
	description = "description"
)

func TestService_CreateNewUser(t *testing.T) {
	ctx := &gin.Context{}
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	newBoard := &model.Board{
		UserID:      userID,
		Title:       title,
		Description: null.StringFrom(description),
	}

	masterTx := repository.NewMockMasterTx()

	boardRepository := mock_board.NewMockRepository(ctrl)
	boardRepository.EXPECT().InsertBoard(ctx, masterTx, newBoard).Return(nil).Times(1)

	service := New(boardRepository)
	err := service.CreateNewBoard(ctx, masterTx, userID, title, description)

	assert.NoError(t, err)
}

func TestService_GetByPK(t *testing.T) {
	ctx := &gin.Context{}
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	existedBoard := &model.Board{
		ID:          id,
		UserID:      userID,
		Title:       title,
		Description: null.StringFrom(description),
	}

	masterTx := repository.NewMockMasterTx()

	boardRepository := mock_board.NewMockRepository(ctrl)
	boardRepository.EXPECT().SelectByPK(ctx, masterTx, id).Return(existedBoard, nil).Times(1)

	service := New(boardRepository)
	boards, err := service.GetByPK(ctx, masterTx, id)

	assert.NoError(t, err)
	assert.NotNil(t, boards)
}

func TestService_GetByUserID(t *testing.T) {
	ctx := &gin.Context{}
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	userBoards := model.BoardSlice{
		{
			ID:          id,
			UserID:      userID,
			Title:       title,
			Description: null.StringFrom(description),
		},
	}

	masterTx := repository.NewMockMasterTx()

	boardRepository := mock_board.NewMockRepository(ctrl)
	boardRepository.EXPECT().SelectByUserID(ctx, masterTx, userID).Return(userBoards, nil).Times(1)

	service := New(boardRepository)
	boards, err := service.GetByUserID(ctx, masterTx, userID)

	assert.NoError(t, err)
	assert.NotNil(t, boards)
}

func TestService_SelectAll(t *testing.T) {
	ctx := &gin.Context{}
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	existedBoards := model.BoardSlice{
		{
			ID:          id,
			UserID:      userID,
			Title:       title,
			Description: null.StringFrom(description),
		},
	}

	masterTx := repository.NewMockMasterTx()

	boardRepository := mock_board.NewMockRepository(ctrl)
	boardRepository.EXPECT().SelectAll(ctx, masterTx).Return(existedBoards, nil).Times(1)

	service := New(boardRepository)
	boards, err := service.GetAll(ctx, masterTx)

	assert.NoError(t, err)
	assert.NotNil(t, boards)
}
