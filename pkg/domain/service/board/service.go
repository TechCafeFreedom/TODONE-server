package board

import (
	"todone/pkg/domain/entity"
	"todone/pkg/domain/repository"
	"todone/pkg/domain/repository/board"

	"github.com/gin-gonic/gin"
)

type Service interface {
	CreateNewBoard(ctx *gin.Context, masterTx repository.MasterTx, userID int, title, description string) error
	GetByPK(ctx *gin.Context, masterTx repository.MasterTx, id int) (*entity.Board, error)
	GetByUserID(ctx *gin.Context, masterTx repository.MasterTx, userID int) (entity.BoardSlice, error)
	GetAll(ctx *gin.Context, masterTx repository.MasterTx) (entity.BoardSlice, error)
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
	if err := s.boardRepository.InsertBoard(ctx, masterTx, userID, title, description); err != nil {
		return err
	}
	return nil
}

func (s *service) GetByPK(ctx *gin.Context, masterTx repository.MasterTx, id int) (*entity.Board, error) {
	boardData, err := s.boardRepository.SelectByPK(ctx, masterTx, id)
	if err != nil {
		return nil, err
	}
	return boardData, nil
}

func (s *service) GetByUserID(ctx *gin.Context, masterTx repository.MasterTx, userID int) (entity.BoardSlice, error) {
	boardSlice, err := s.boardRepository.SelectByUserID(ctx, masterTx, userID)
	if err != nil {
		return nil, err
	}
	return boardSlice, nil
}

func (s *service) GetAll(ctx *gin.Context, masterTx repository.MasterTx) (entity.BoardSlice, error) {
	boardSlice, err := s.boardRepository.SelectAll(ctx, masterTx)
	if err != nil {
		return nil, err
	}
	return boardSlice, nil
}
