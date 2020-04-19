package user

import (
	"todone/pkg/domain/entity"
	"todone/pkg/domain/repository"

	"github.com/gin-gonic/gin"
)

type Repository interface {
	InsertUser(ctx *gin.Context, masterTx repository.MasterTx, uid, name, thumbnail string) error
	SelectByPK(ctx *gin.Context, masterTx repository.MasterTx, userID int) (*entity.User, error)
	SelectByUID(ctx *gin.Context, masterTx repository.MasterTx, uid string) (*entity.User, error)
	SelectAll(ctx *gin.Context, masterTx repository.MasterTx) (entity.UserSlice, error)
}
