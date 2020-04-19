package board

import (
	"testing"
	"todone/pkg/domain/entity"
	"todone/pkg/domain/repository"
	"todone/pkg/domain/repository/board/mock_board"

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

	masterTx := repository.NewMockMasterTx()

	boardRepository := mock_board.NewMockRepository(ctrl)
	boardRepository.EXPECT().InsertBoard(ctx, masterTx, userID, title, description).Return(nil).Times(1)

	service := New(boardRepository)
	err := service.CreateNewBoard(ctx, masterTx, userID, title, description)

	assert.NoError(t, err)
}

func TestService_GetByPK(t *testing.T) {
	ctx := &gin.Context{}
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	existedBoard := &entity.Board{
		ID:          id,
		Title:       title,
		Description: description,
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

	userBoards := entity.BoardSlice{
		{
			ID:          id,
			Title:       title,
			Description: description,
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

	existedBoards := entity.BoardSlice{
		{
			ID:          id,
			Title:       title,
			Description: description,
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
