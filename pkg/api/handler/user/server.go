package user

import (
	"net/http"
	"todone/pkg/api/middleware"
	"todone/pkg/api/request/reqbody"
	"todone/pkg/api/response"
	userinteractor "todone/pkg/api/usecase/user"
	"todone/pkg/terrors"

	"github.com/gin-gonic/gin"
)

type Server struct {
	userInteractor userinteractor.Interactor
}

func New(userInteractor userinteractor.Interactor) Server {
	return Server{userInteractor: userInteractor}
}

func (s *Server) CreateNewUser(ctx *gin.Context) {
	var reqBody reqbody.UserCreate
	if err := ctx.BindJSON(&reqBody); err != nil {
		ctx.Error(terrors.Stack(err))
		return
	}

	uid, ok := ctx.Get(middleware.AuthCtxKey)
	if !ok {
		errMessageJP := "不正なユーザからのアクセスをブロックしました。"
		errMessageEN := "The content blocked because user is not certified."
		ctx.Error(terrors.Newf(http.StatusUnauthorized, errMessageJP, errMessageEN))
		return
	}

	if err := s.userInteractor.CreateNewUser(ctx, uid.(string), reqBody.Name, reqBody.Thumbnail); err != nil {
		ctx.Error(terrors.Stack(err))
		return
	}

	ctx.Status(http.StatusNoContent)
	return
}

func (s *Server) GetUserProfile(ctx *gin.Context) {
	uid, ok := ctx.Get(middleware.AuthCtxKey)
	if !ok {
		errMessageJP := "不正なユーザからのアクセスをブロックしました。"
		errMessageEN := "The content blocked because user is not certified."
		ctx.Error(terrors.Newf(http.StatusUnauthorized, errMessageJP, errMessageEN))
		return
	}

	user, err := s.userInteractor.GetUserProfile(ctx, uid.(string))
	if err != nil {
		ctx.Error(terrors.Stack(err))
		return
	}

	ctx.JSON(http.StatusOK, response.ConvertToUserResponse(user))
	return
}

func (s *Server) GetAllUsers(ctx *gin.Context) {
	users, err := s.userInteractor.GetAll(ctx)
	if err != nil {
		ctx.Error(terrors.Stack(err))
		return
	}

	ctx.JSON(http.StatusOK, response.ConvertToUsersResponse(users))
	return
}
