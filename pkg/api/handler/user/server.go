package user

import (
	"errors"
	"net/http"
	"todone/pkg/api/middleware"
	"todone/pkg/api/request/reqbody"
	"todone/pkg/api/response"
	userinteractor "todone/pkg/api/usecase/user"

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
		ctx.Error(err)
	}

	uid, ok := ctx.Get(middleware.AuthCtxKey)
	if !ok {
		ctx.Error(errors.New("access token is not found in context"))
	}

	if err := s.userInteractor.CreateNewUser(ctx, uid.(string), reqBody.Name, reqBody.Thumbnail); err != nil {
		ctx.Error(err)
	}

	ctx.Status(http.StatusNoContent)
}

func (s *Server) GetUserProfile(ctx *gin.Context) {
	uid, ok := ctx.Get(middleware.AuthCtxKey)
	if !ok {
		ctx.Error(errors.New("access token is not found in context"))
	}

	user, err := s.userInteractor.GetUserProfile(ctx, uid.(string))
	if err != nil {
		ctx.Error(err)
	}

	ctx.JSON(http.StatusOK, response.ConvertToUserResponse(user))
}

func (s *Server) GetAllUsers(ctx *gin.Context) {
	users, err := s.userInteractor.GetAll(ctx)
	if err != nil {
		ctx.Error(err)
	}

	ctx.JSON(http.StatusOK, response.ConvertToUsersResponse(users))
}
