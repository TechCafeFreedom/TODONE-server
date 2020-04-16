package board

import (
	"todone/db/mysql/model"
	"todone/pkg/domain/repository"

	"github.com/gin-gonic/gin"
)

type Repository interface {
	InsertBoard(ctx *gin.Context, masterTx repository.MasterTx, boardData *model.Board) error
	SelectByPK(ctx *gin.Context, masterTx repository.MasterTx, id int) (*model.Board, error)
	SelectByUserID(ctx *gin.Context, masterTx repository.MasterTx, userID int) (model.BoardSlice, error)
	SelectAll(ctx *gin.Context, masterTx repository.MasterTx) (model.BoardSlice, error)
}
