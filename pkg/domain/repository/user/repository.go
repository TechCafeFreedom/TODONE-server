package user

import (
	"todone/db/mysql/model"
	"todone/pkg/infrastructure/mysql"

	"github.com/gin-gonic/gin"
)

type Repository interface {
	InsertUser(ctx *gin.Context, masterTx mysql.MasterTx, user *model.User) error
	SelectByPK(ctx *gin.Context, masterTx mysql.MasterTx, userID string) (*model.User, error)
	SelectAll(ctx *gin.Context, masterTx mysql.MasterTx) (model.UserSlice, error)
}
