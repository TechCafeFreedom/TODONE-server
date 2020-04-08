package project

import (
	"todone/db/mysql/model"
	"todone/pkg/infrastructure/mysql"

	"github.com/gin-gonic/gin"
)

type Repository interface {
	InsertProject(ctx *gin.Context, masterTx mysql.MasterTx, project *model.Project) error
	SelectByPK(ctx *gin.Context, masterTx mysql.MasterTx, id int) (*model.Project, error)
	SelectByUserID(ctx *gin.Context, masterTx mysql.MasterTx, userID string) (model.ProjectSlice, error)
	SelectAll(ctx *gin.Context, masterTx mysql.MasterTx) (model.ProjectSlice, error)
}
