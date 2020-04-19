package board

import (
	"errors"
	"net/http"
	"strconv"
	"todone/pkg/api/middleware"
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

	uid, ok := ctx.Get(middleware.AuthCtxKey)
	if !ok {
		ctx.Error(errors.New("access token is not found in context"))
	}

	if err := s.boardInteractor.CreateNewBoard(ctx, uid.(string), reqBody.Title, reqBody.Description); err != nil {
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

	ctx.JSON(http.StatusOK, response.ConvertToBoardResponse(board))
}

func (s *Server) GetUserBoards(ctx *gin.Context) {
	uid, ok := ctx.Get(middleware.AuthCtxKey)
	if !ok {
		ctx.Error(errors.New("access token is not found in context"))
	}

	boards, err := s.boardInteractor.GetUserBoards(ctx, uid.(string))
	if err != nil {
		ctx.Error(err)
	}

	ctx.JSON(http.StatusOK, response.ConvertToBoardsResponse(boards))
}

func (s *Server) GetAllBoards(ctx *gin.Context) {
	boards, err := s.boardInteractor.GetAll(ctx)
	if err != nil {
		ctx.Error(err)
	}

	ctx.JSON(http.StatusOK, response.ConvertToBoardsResponse(boards))
}
