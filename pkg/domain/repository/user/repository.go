package user

import (
	"todone/db/mysql/model"
	"todone/pkg/domain/entity"
	"todone/pkg/domain/repository"

	"github.com/gin-gonic/gin"
)

type Repository interface {
	InsertUser(ctx *gin.Context, masterTx repository.MasterTx, userData *model.User) error
	SelectByPK(ctx *gin.Context, masterTx repository.MasterTx, userID int) (*entity.User, error)
	SelectByAccessToken(ctx *gin.Context, masterTx repository.MasterTx, accessToken string) (*entity.User, error)
	SelectAll(ctx *gin.Context, masterTx repository.MasterTx) (entity.UserSlice, error)
}
