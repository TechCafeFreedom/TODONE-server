package board

import (
	"context"
	"net/http"
	"todone/pkg/domain/entity"
	"todone/pkg/domain/repository"
	"todone/pkg/domain/repository/board"
	"todone/pkg/terrors"
)

type Service interface {
	CreateNewBoard(ctx context.Context, masterTx repository.MasterTx, userID int, title, description string) error
	GetByPK(ctx context.Context, masterTx repository.MasterTx, id int) (*entity.Board, error)
	GetByUserID(ctx context.Context, masterTx repository.MasterTx, userID int) (entity.BoardSlice, error)
	GetAll(ctx context.Context, masterTx repository.MasterTx) (entity.BoardSlice, error)
}

type service struct {
	boardRepository board.Repository
}

func New(boardRepository board.Repository) Service {
	return &service{
		boardRepository: boardRepository,
	}
}

func (s *service) CreateNewBoard(ctx context.Context, masterTx repository.MasterTx, userID int, title, description string) error {
	// titleの空文字チェック
	if title == "" {
		return terrors.Newf(http.StatusBadRequest, "ボードタイトルは必須項目です。", "Board title is required.")
	}

	if err := s.boardRepository.InsertBoard(ctx, masterTx, userID, title, description); err != nil {
		return terrors.Stack(err)
	}
	return nil
}

func (s *service) GetByPK(ctx context.Context, masterTx repository.MasterTx, id int) (*entity.Board, error) {
	boardData, err := s.boardRepository.SelectByPK(ctx, masterTx, id)
	if err != nil {
		return nil, terrors.Stack(err)
	}
	return boardData, nil
}

func (s *service) GetByUserID(ctx context.Context, masterTx repository.MasterTx, userID int) (entity.BoardSlice, error) {
	boardSlice, err := s.boardRepository.SelectByUserID(ctx, masterTx, userID)
	if err != nil {
		return nil, terrors.Stack(err)
	}
	return boardSlice, nil
}

func (s *service) GetAll(ctx context.Context, masterTx repository.MasterTx) (entity.BoardSlice, error) {
	boardSlice, err := s.boardRepository.SelectAll(ctx, masterTx)
	if err != nil {
		return nil, terrors.Stack(err)
	}
	return boardSlice, nil
}
