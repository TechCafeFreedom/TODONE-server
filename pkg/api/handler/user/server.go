package user

import (
	"errors"
	"net/http"
	"todone/db/mysql/model"
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

	accessToken, ok := ctx.Get("AUTHED_ACCESS_TOKEN")
	if !ok {
		ctx.Error(errors.New("userID is not found in context"))
	}

	if err := s.userInteractor.CreateNewUser(ctx, accessToken.(string), reqBody.Name, reqBody.Thumbnail); err != nil {
		ctx.Error(err)
	}

	ctx.Status(http.StatusNoContent)
}

func (s *Server) GetUserProfile(ctx *gin.Context) {
	userID, ok := ctx.Get("AUTHED_ACCESS_TOKEN")
	if !ok {
		ctx.Error(errors.New("userID is not found in context"))
	}

	user, err := s.userInteractor.GetUserProfile(ctx, userID.(string))
	if err != nil {
		ctx.Error(err)
	}

	ctx.JSON(http.StatusOK, convertToUserResponse(user))
}

func (s *Server) GetAllUsers(ctx *gin.Context) {
	users, err := s.userInteractor.GetAll(ctx)
	if err != nil {
		ctx.Error(err)
	}

	ctx.JSON(http.StatusOK, convertToUsersResponse(users))
}

func convertToUserResponse(user *model.User) response.UserResponse {
	return response.UserResponse{
		Name:      user.Name,
		Thumbnail: user.Thumbnail.String,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}

func convertToUsersResponse(users model.UserSlice) response.UsersResponse {
	res := make(response.UsersResponse, 0, len(users))
	for _, data := range users {
		res = append(res, &response.UserResponse{
			Name:      data.Name,
			Thumbnail: data.Thumbnail.String,
			CreatedAt: data.CreatedAt,
			UpdatedAt: data.UpdatedAt,
		})
	}
	return res
}
