package user

import (
	"net/http"
	"os"
	"time"
	"todone/pkg/api/reqbody"
	"todone/pkg/domain/service/user"
	"todone/pkg/utility"

	ginJwt "github.com/appleboy/gin-jwt/v2"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type Server struct {
	UserService user.Service
}

func NewUserHandler(service user.Service) Server {
	return Server{UserService: service}
}

func (handler *Server) Login(context *gin.Context) (interface{}, error) {
	var reqBody reqbody.UserLogin

	err := context.BindJSON(&reqBody)
	if err != nil {
		return nil, err
	}

	loginedUser, err := handler.UserService.Login(reqBody.Email, reqBody.Password)
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

	user, err := handler.UserService.GetUser(id)
	if err != nil {
		context.Error(err)
	}

	context.JSON(200, user)
}

func (handler *Server) CreateNewUser(context *gin.Context) {
	var reqBody reqbody.UserCreate

	err := context.BindJSON(&reqBody)
	if err != nil {
		context.Error(err)
	}

	id := utility.CreateUserID(255)

	err = handler.UserService.CreateNewUser(id, "userID", reqBody.Name, reqBody.Thumbnail, reqBody.Email, reqBody.Password)
	if err != nil {
		context.Error(err)
	}

	// token作成処理
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"exp":      time.Now().Add(time.Hour).Unix(),
		"id":       id,
		"orig_iat": time.Now().Unix(),
	})
	tokenStr, err := token.SignedString([]byte(os.Getenv("JWT_KEY")))
	if err != nil {
		context.Error(err)
	}

	context.JSON(http.StatusOK, gin.H{"token": tokenStr})
}

func (handler *Server) GetAllUsers(context *gin.Context) {
	users, err := handler.UserService.SelectAll()
	if err != nil {
		context.Error(err)
	}

	context.JSON(http.StatusOK, gin.H{"testUsers": users})
}
