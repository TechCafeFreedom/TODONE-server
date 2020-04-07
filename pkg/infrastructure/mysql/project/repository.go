package project

import (
	"context"
	"todone/db/mysql/model"
	"todone/pkg/domain/repository/project"

	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries/qm"
)

type projectRepositoryImpliment struct {
	db boil.ContextExecutor
}

func New(db boil.ContextExecutor) project.Repository {
	return &projectRepositoryImpliment{
		db: db,
	}
}

// プロジェクト作成機能
func (p *projectRepositoryImpliment) InsertProject(project *model.Project) error {
	if err := project.Insert(context.Background(), p.db, boil.Infer()); err != nil {
		return err
	}

	return nil
}

// プロジェクト取得機能（PK: id）
func (p *projectRepositoryImpliment) SelectByPK(id int) (*model.Project, error) {
	project, err := model.FindProject(context.Background(), p.db, id)
	if err != nil {
		return nil, err
	}

	return project, nil
}

// ユーザのもつプロジェクト取得機能
func (p *projectRepositoryImpliment) SelectByUserID(userID string) (model.ProjectSlice, error) {
	projects, err := model.Projects(model.ProjectWhere.UserID.EQ(userID)).All(context.Background(), p.db)
	if err != nil {
		return nil, err
	}

	return projects, nil
}

// プロジェクト全件取得機能
func (p *projectRepositoryImpliment) SelectAll() (model.ProjectSlice, error) {
	queries := []qm.QueryMod{}
	projects, err := model.Projects(queries...).All(context.Background(), p.db)
	if err != nil {
		return nil, err
	}

	return projects, nil
}
