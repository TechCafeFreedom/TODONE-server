package user

import (
	"todone/db/mysql/model"
	"todone/pkg/domain/repository"
	"todone/pkg/domain/repository/user"

	"github.com/gin-gonic/gin"
	"github.com/volatiletech/null"
)

type Service interface {
	CreateNewUser(ctx *gin.Context, masterTx repository.MasterTx, accessToken, name, thumbnail string) error
	GetByPK(ctx *gin.Context, masterTx repository.MasterTx, userID int) (*model.User, error)
	GetByAccessToken(ctx *gin.Context, masterTx repository.MasterTx, accessToken string) (*model.User, error)
	GetAll(ctx *gin.Context, masterTx repository.MasterTx) (model.UserSlice, error)
}

type service struct {
	userRepository user.Repository
}

func New(userRepository user.Repository) Service {
	return &service{
		userRepository: userRepository,
	}
}

func (s *service) CreateNewUser(ctx *gin.Context, masterTx repository.MasterTx, accessToken, name, thumbnail string) error {
	if err := s.userRepository.InsertUser(ctx, masterTx, &model.User{
		AccessToken: accessToken,
		Name:        name,
		Thumbnail:   null.StringFrom(thumbnail),
	}); err != nil {
		return err
	}
	return nil
}

func (s *service) GetByPK(ctx *gin.Context, masterTx repository.MasterTx, userID int) (*model.User, error) {
	userData, err := s.userRepository.SelectByPK(ctx, masterTx, userID)
	if err != nil {
		return nil, err
	}
	return userData, nil
}

func (s *service) GetByAccessToken(ctx *gin.Context, masterTx repository.MasterTx, accessToken string) (*model.User, error) {
	userData, err := s.userRepository.SelectByAccessToken(ctx, masterTx, accessToken)
	if err != nil {
		return nil, err
	}
	return userData, nil
}

func (s *service) GetAll(ctx *gin.Context, masterTx repository.MasterTx) (model.UserSlice, error) {
	return s.userRepository.SelectAll(ctx, masterTx)
}
