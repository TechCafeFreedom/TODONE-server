package board

import (
	"todone/pkg/domain/entity"
	"todone/pkg/domain/repository"

	"github.com/gin-gonic/gin"
)

type Repository interface {
	InsertBoard(ctx *gin.Context, masterTx repository.MasterTx, userID int, title, description string) error
	SelectByPK(ctx *gin.Context, masterTx repository.MasterTx, id int) (*entity.Board, error)
	SelectByUserID(ctx *gin.Context, masterTx repository.MasterTx, userID int) (entity.BoardSlice, error)
	SelectAll(ctx *gin.Context, masterTx repository.MasterTx) (entity.BoardSlice, error)
}
