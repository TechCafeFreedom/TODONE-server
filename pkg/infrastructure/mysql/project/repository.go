package project

import (
	"todone/db/mysql/model"
	"todone/pkg/domain/repository"
	"todone/pkg/domain/repository/project"
	"todone/pkg/infrastructure/mysql"

	"github.com/gin-gonic/gin"
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries/qm"
)

type projectRepositoryImpliment struct {
	masterTxManager repository.MasterTxManager
}

func New(masterTxManager repository.MasterTxManager) project.Repository {
	return &projectRepositoryImpliment{
		masterTxManager: masterTxManager,
	}
}

// プロジェクト作成機能
func (p *projectRepositoryImpliment) InsertProject(ctx *gin.Context, masterTx repository.MasterTx, projectData *model.Project) error {
	exec, err := mysql.ExtractExecutor(masterTx)
	if err != nil {
		return err
	}
	if err := projectData.Insert(ctx, exec, boil.Infer()); err != nil {
		return err
	}

	return nil
}

// プロジェクト取得機能（PK: id）
func (p *projectRepositoryImpliment) SelectByPK(ctx *gin.Context, masterTx repository.MasterTx, id int) (*model.Project, error) {
	exec, err := mysql.ExtractExecutor(masterTx)
	if err != nil {
		return nil, err
	}
	projectData, err := model.FindProject(ctx, exec, id)
	if err != nil {
		return nil, err
	}

	return projectData, nil
}

// ユーザのもつプロジェクト取得機能
func (p *projectRepositoryImpliment) SelectByUserID(ctx *gin.Context, masterTx repository.MasterTx, userID string) (model.ProjectSlice, error) {
	exec, err := mysql.ExtractExecutor(masterTx)
	if err != nil {
		return nil, err
	}
	projects, err := model.Projects(model.ProjectWhere.UserID.EQ(userID)).All(ctx, exec)
	if err != nil {
		return nil, err
	}

	return projects, nil
}

// プロジェクト全件取得機能
func (p *projectRepositoryImpliment) SelectAll(ctx *gin.Context, masterTx repository.MasterTx) (model.ProjectSlice, error) {
	exec, err := mysql.ExtractExecutor(masterTx)
	if err != nil {
		return nil, err
	}
	queries := []qm.QueryMod{}
	projects, err := model.Projects(queries...).All(ctx, exec)
	if err != nil {
		return nil, err
	}

	return projects, nil
}
