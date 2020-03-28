package user

import (
	model "todone/db/mysql/models"
	"todone/pkg/domain/repository/user"

	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	Login(email, password string) (*model.User, error)
	GetUser(id string) ([]*model.User, error)
	CreateNewUser(id, userID, name, thumbnail, email, password string) error
	SelectAll() ([]*model.User, error)
}

type service struct {
	userRepository user.Repository
}

func NewUserService(repository user.Repository) Service {
	return &service{
		userRepository: repository,
	}
}

func (s *service) Login(email, password string) (*model.User, error) {
	return s.userRepository.Login(email, password)
}

func (s *service) GetUser(id string) ([]*model.User, error) {
	return s.userRepository.SelectByUserId(id)
}

func (s *service) CreateNewUser(id, userID, name, thumbnail, email, password string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user := &model.User{
		ID:        id,
		UserID:    userID,
		Email:     email,
		Password:  string(hashedPassword),
		Name:      name,
		Thumbnail: thumbnail,
	}

	return s.userRepository.InsertUser(user)

}

func (s *service) SelectAll() ([]*model.User, error) {
	return s.userRepository.SelectAll()
}
