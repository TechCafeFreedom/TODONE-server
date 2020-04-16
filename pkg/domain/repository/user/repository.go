package user

import (
	"todone/db/mysql/model"
	"todone/pkg/domain/repository"

	"github.com/gin-gonic/gin"
)

type Repository interface {
	InsertUser(ctx *gin.Context, masterTx repository.MasterTx, userData *model.User) error
	SelectByPK(ctx *gin.Context, masterTx repository.MasterTx, userID int) (*model.User, error)
	SelectByAccessToken(ctx *gin.Context, masterTx repository.MasterTx, accessToken string) (*model.User, error)
	SelectAll(ctx *gin.Context, masterTx repository.MasterTx) (model.UserSlice, error)
}
