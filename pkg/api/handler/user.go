package handler

import (
	"net/http"
	"os"
	"time"
	"todone/pkg/api/requestBody"
	"todone/pkg/domain/service/user"
	"todone/pkg/utility"

	ginJwt "github.com/appleboy/gin-jwt/v2"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	UserService user.Service
}

func NewUserHandler(service user.Service) UserHandler {
	return UserHandler{UserService: service}
}

func (handler *UserHandler) Login(context *gin.Context) (interface{}, error) {
	var reqBody requestBody.UserLogin

	err := context.BindJSON(&reqBody)
	if err != nil {
		return "", err
	}

	user, err := handler.UserService.Login(reqBody.Email, reqBody.Password)
	if err != nil {
		return nil, err
	} else if user == nil {
		return nil, ginJwt.ErrFailedAuthentication
	}

	return user, nil
}

func (handler *UserHandler) GetUser(context *gin.Context) {
	claims := ginJwt.ExtractClaims(context)
	id, ok := claims["id"].(string)
	if ok == false {
		context.Error(ginJwt.ErrForbidden)
	}

	user, err := handler.UserService.GetUser(id)
	if err != nil {
		context.Error(err)
	}

	context.JSON(200, user)
}

func (handler *UserHandler) CreateNewUser(context *gin.Context) {
	var reqBody requestBody.UserCreate

	err := context.BindJSON(&reqBody)
	if err != nil {
		context.Error(err)
	}

	id := utility.CreateUserId(255)

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

func (handler *UserHandler) GetAllUsers(context *gin.Context) {
	users, err := handler.UserService.SelectAll()
	if err != nil {
		context.Error(err)
	}

	context.JSON(http.StatusOK, gin.H{"testUsers": users})
}
