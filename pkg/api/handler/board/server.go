package board

import (
	"fmt"
	"net/http"
	"strconv"
	"todone/pkg/api/middleware"
	"todone/pkg/api/request/reqbody"
	"todone/pkg/api/response"
	boardinteractor "todone/pkg/api/usecase/board"
	"todone/pkg/terrors"

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
		ctx.Error(terrors.Stack(err))
		return
	}

	uid, ok := ctx.Get(middleware.AuthCtxKey)
	if !ok {
		errMessageJP := fmt.Sprintf("uidはリクエストされたコンテキストから検出されませんでした。ctx: %v", ctx)
		errMessageEN := fmt.Sprintf("uid is not found in context. ctx: %v", ctx)
		ctx.Error(terrors.Newf(http.StatusUnauthorized, errMessageJP, errMessageEN))
		return
	}

	if err := s.boardInteractor.CreateNewBoard(ctx, uid.(string), reqBody.Title, reqBody.Description); err != nil {
		ctx.Error(err)
		return
	}

	ctx.Status(http.StatusNoContent)
	return
}

func (s *Server) GetBoardDetail(ctx *gin.Context) {
	idParam := ctx.Param("id")
	boardID, err := strconv.Atoi(idParam)
	if err != nil {
		ctx.Error(terrors.Stack(err))
		return
	}

	board, err := s.boardInteractor.GetBoardDetail(ctx, boardID)
	if err != nil {
		ctx.Error(terrors.Stack(err))
		return
	}

	ctx.JSON(http.StatusOK, response.ConvertToBoardResponse(board))
	return
}

func (s *Server) GetUserBoards(ctx *gin.Context) {
	uid, ok := ctx.Get(middleware.AuthCtxKey)
	if !ok {
		errMessageJP := fmt.Sprintf("uidはリクエストされたコンテキストから検出されませんでした。ctx: %v", ctx)
		errMessageEN := fmt.Sprintf("uid is not found in context. ctx: %v", ctx)
		ctx.Error(terrors.Newf(http.StatusUnauthorized, errMessageJP, errMessageEN))
		return
	}

	boards, err := s.boardInteractor.GetUserBoards(ctx, uid.(string))
	if err != nil {
		ctx.Error(terrors.Stack(err))
		return
	}

	ctx.JSON(http.StatusOK, response.ConvertToBoardsResponse(boards))
}

func (s *Server) GetAllBoards(ctx *gin.Context) {
	boards, err := s.boardInteractor.GetAll(ctx)
	if err != nil {
		ctx.Error(terrors.Stack(err))
		return
	}

	ctx.JSON(http.StatusOK, response.ConvertToBoardsResponse(boards))
	return
}
