package board

import (
	"todone/db/mysql/model"
	"todone/pkg/domain/repository"
	"todone/pkg/domain/repository/board"

	"github.com/volatiletech/null"

	"github.com/gin-gonic/gin"
)

type Service interface {
	CreateNewBoard(ctx *gin.Context, masterTx repository.MasterTx, userID int, title, description string) error
	GetByPK(ctx *gin.Context, masterTx repository.MasterTx, id int) (*model.Board, error)
	GetByUserID(ctx *gin.Context, masterTx repository.MasterTx, userID int) (model.BoardSlice, error)
	GetAll(ctx *gin.Context, masterTx repository.MasterTx) (model.BoardSlice, error)
}

type service struct {
	boardRepository board.Repository
}

func New(boardRepository board.Repository) Service {
	return &service{
		boardRepository: boardRepository,
	}
}

func (s *service) CreateNewBoard(ctx *gin.Context, masterTx repository.MasterTx, userID int, title, description string) error {
	if err := s.boardRepository.InsertBoard(ctx, masterTx, &model.Board{
		UserID:      userID,
		Title:       title,
		Description: null.StringFrom(description),
	}); err != nil {
		return err
	}
	return nil
}

func (s *service) GetByPK(ctx *gin.Context, masterTx repository.MasterTx, id int) (*model.Board, error) {
	boardData, err := s.boardRepository.SelectByPK(ctx, masterTx, id)
	if err != nil {
		return nil, err
	}
	return boardData, nil
}

func (s *service) GetByUserID(ctx *gin.Context, masterTx repository.MasterTx, userID int) (model.BoardSlice, error) {
	boards, err := s.boardRepository.SelectByUserID(ctx, masterTx, userID)
	if err != nil {
		return nil, err
	}
	return boards, nil
}

func (s *service) GetAll(ctx *gin.Context, masterTx repository.MasterTx) (model.BoardSlice, error) {
	return s.boardRepository.SelectAll(ctx, masterTx)
}
