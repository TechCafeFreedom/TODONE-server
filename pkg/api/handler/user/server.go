package user

import (
	"net/http"
	"todone/pkg/api/reqbody"
	userinteractor "todone/pkg/api/usecase/user"

	ginJwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

type Server struct {
	userInteractor userinteractor.Interactor
}

func New(userInteractor userinteractor.Interactor) Server {
	return Server{userInteractor: userInteractor}
}

func (handler *Server) Login(context *gin.Context) (interface{}, error) {
	var reqBody reqbody.UserLogin

	err := context.BindJSON(&reqBody)
	if err != nil {
		return nil, err
	}

	loginedUser, err := handler.userInteractor.Login(reqBody.Email, reqBody.Password)
	if err != nil {
		return nil, err
	} else if loginedUser == nil {
		return nil, ginJwt.ErrFailedAuthentication
	}

	return loginedUser, nil
}

func (handler *Server) GetUser(context *gin.Context) {
	claims := ginJwt.ExtractClaims(context)
	id, ok := claims["id"].(string)
	if !ok {
		context.Error(ginJwt.ErrForbidden)
	}

	user, err := handler.userInteractor.GetUser(id)
	if err != nil {
		context.Error(err)
	}

	context.JSON(http.StatusOK, user)
}

func (handler *Server) CreateNewUser(context *gin.Context) {
	var reqBody reqbody.UserCreate

	err := context.BindJSON(&reqBody)
	if err != nil {
		context.Error(err)
	}

	tokenStr, err := handler.userInteractor.CreateNewUser("userID", reqBody.Name, reqBody.Thumbnail, reqBody.Email, reqBody.Password)
	if err != nil {
		context.Error(err)
	}

	context.JSON(http.StatusOK, gin.H{"token": tokenStr})
}

func (handler *Server) GetAllUsers(context *gin.Context) {
	users, err := handler.userInteractor.SelectAll()
	if err != nil {
		context.Error(err)
	}

	context.JSON(http.StatusOK, gin.H{"testUsers": users})
}
