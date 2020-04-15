package project

import (
	"todone/db/mysql/model"
	"todone/pkg/domain/repository"

	"github.com/gin-gonic/gin"
)

type Repository interface {
	InsertProject(ctx *gin.Context, masterTx repository.MasterTx, projectData *model.Project) error
	SelectByPK(ctx *gin.Context, masterTx repository.MasterTx, id int) (*model.Project, error)
	SelectByUserID(ctx *gin.Context, masterTx repository.MasterTx, userID string) (model.ProjectSlice, error)
	SelectAll(ctx *gin.Context, masterTx repository.MasterTx) (model.ProjectSlice, error)
}
