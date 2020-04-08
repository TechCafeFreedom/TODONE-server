package user

import (
	"todone/db/mysql/model"
	"todone/pkg/domain/repository/user"

	"github.com/volatiletech/null"
)

type Service interface {
	CreateNewUser(userID, name, thumbnail string) error
	GetByPK(userID string) (*model.User, error)
	GetAll() (model.UserSlice, error)
}

type service struct {
	userRepository user.Repository
}

func New(userRepository user.Repository) Service {
	return &service{
		userRepository: userRepository,
	}
}

func (s *service) CreateNewUser(userID, name, thumbnail string) error {
	if err := s.userRepository.InsertUser(&model.User{
		UserID:    userID,
		Name:      name,
		Thumbnail: null.StringFrom(thumbnail),
	}); err != nil {
		return err
	}
	return nil
}

func (s *service) GetByPK(userID string) (*model.User, error) {
	userData, err := s.userRepository.SelectByPK(userID)
	if err != nil {
		return nil, err
	}
	return userData, nil
}

func (s *service) GetAll() (model.UserSlice, error) {
	return s.userRepository.SelectAll()
}
