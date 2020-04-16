package board

import (
	"errors"
	"net/http"
	"strconv"
	"todone/db/mysql/model"
	"todone/pkg/api/request/reqbody"
	"todone/pkg/api/response"
	boardinteractor "todone/pkg/api/usecase/board"

	"github.com/gin-gonic/gin"
)

type Server struct {
	boardInteractor boardinteractor.Interactor
}

func New(boardInteractor boardinteractor.Interactor) Server {
	return Server{boardInteractor: boardInteractor}
}

func (s *Server) CreateNewBoard(ctx *gin.Context) {
	var reqBody reqbody.BoardCreate
	if err := ctx.BindJSON(&reqBody); err != nil {
		ctx.Error(err)
	}

	accessToken, ok := ctx.Get("AUTHED_ACCESS_TOKEN")
	if !ok {
		ctx.Error(errors.New("userID is not found in context"))
	}

	if err := s.boardInteractor.CreateNewBoard(ctx, accessToken.(string), reqBody.Title, reqBody.Description); err != nil {
		ctx.Error(err)
	}

	ctx.Status(http.StatusNoContent)
}

func (s *Server) GetBoardDetail(ctx *gin.Context) {
	idParam := ctx.Param("id")
	boardID, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.Error(err)
	}

	board, err := s.boardInteractor.GetBoardDetail(ctx, boardID)
	if err != nil {
		ctx.Error(err)
	}

	ctx.JSON(http.StatusOK, convertToBoardResponse(board))
}

func (s *Server) GetUserBoards(ctx *gin.Context) {
	accessToken, ok := ctx.Get("AUTHED_ACCESS_TOKEN")
	if !ok {
		ctx.Error(errors.New("userID is not found in context"))
	}

	boards, err := s.boardInteractor.GetUserBoards(ctx, accessToken.(string))
	if err != nil {
		ctx.Error(err)
	}

	ctx.JSON(http.StatusOK, convertToBoardsResponse(boards))
}

func (s *Server) GetAllBoards(ctx *gin.Context) {
	boards, err := s.boardInteractor.GetAll(ctx)
	if err != nil {
		ctx.Error(err)
	}

	ctx.JSON(http.StatusOK, convertToBoardsResponse(boards))
}

func convertToBoardResponse(board *model.Board) response.BoardResponse {
	return response.BoardResponse{
		ID:          board.ID,
		Title:       board.Title,
		Description: board.Description.String,
		CreatedAt:   board.CreatedAt,
		UpdatedAt:   board.UpdatedAt,
	}
}

func convertToBoardsResponse(boards model.BoardSlice) response.BoardsResponse {
	res := make(response.BoardsResponse, 0, len(boards))
	for _, data := range boards {
		res = append(res, &response.BoardResponse{
			ID:          data.ID,
			Title:       data.Title,
			Description: data.Description.String,
			CreatedAt:   data.CreatedAt,
			UpdatedAt:   data.UpdatedAt,
		})
	}
	return res
}
