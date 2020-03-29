package user

import (
	"os"
	"time"
	model "todone/db/mysql/models"
	userservice "todone/pkg/domain/service/user"
	"todone/pkg/utility"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

type Interactor interface {
	Login(email, password string) (*model.User, error)
	GetUser(id string) ([]*model.User, error)
	CreateNewUser(userID, name, thumbnail, email, password string) (string, error)
	SelectAll() ([]*model.User, error)
}

type intereractor struct {
	userService userservice.Service
}

func New(userService userservice.Service) Interactor {
	return &intereractor{
		userService: userService,
	}
}

func (i *intereractor) Login(email, password string) (*model.User, error) {
	loginedUser, err := i.userService.Login(email, password)
	if err != nil {
		return nil, err
	}
	return loginedUser, nil
}

func (i *intereractor) GetUser(id string) ([]*model.User, error) {
	users, err := i.userService.GetUser(id)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (i *intereractor) CreateNewUser(userID, name, thumbnail, email, password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	id := utility.CreateUserID(255)

	if err := i.userService.CreateNewUser(&model.User{
		ID:        id,
		UserID:    userID,
		Email:     email,
		Password:  string(hashedPassword),
		Name:      name,
		Thumbnail: thumbnail,
	}); err != nil {
		return "", err
	}

	// token作成処理
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"exp":      time.Now().Add(time.Hour).Unix(),
		"id":       id,
		"orig_iat": time.Now().Unix(),
	})
	tokenStr, err := token.SignedString([]byte(os.Getenv("JWT_KEY")))
	if err != nil {
		return "", err
	}
	return tokenStr, nil
}

func (i *intereractor) SelectAll() ([]*model.User, error) {
	users, err := i.userService.SelectAll()
	if err != nil {
		return nil, err
	}
	return users, nil
}
