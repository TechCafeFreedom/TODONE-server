package user

import (
	model "todone/db/mysql/models"
	"todone/pkg/domain/repository/user"
)

type Service interface {
	Login(email, password string) (*model.User, error)
	GetUser(id string) ([]*model.User, error)
	CreateNewUser(newUser *model.User) error
	SelectAll() ([]*model.User, error)
}

type service struct {
	userRepository user.Repository
}

func New(userRepository user.Repository) Service {
	return &service{
		userRepository: userRepository,
	}
}

func (s *service) Login(email, password string) (*model.User, error) {
	return s.userRepository.Login(email, password)
}

func (s *service) GetUser(id string) ([]*model.User, error) {
	return s.userRepository.SelectByUserID(id)
}

func (s *service) CreateNewUser(newUser *model.User) error {
	return s.userRepository.InsertUser(newUser)
}

func (s *service) SelectAll() ([]*model.User, error) {
	return s.userRepository.SelectAll()
}
