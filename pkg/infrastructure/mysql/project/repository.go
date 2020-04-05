package project

import (
	"context"
	model "todone/db/mysql/models"
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

// プロジェクト全件取得
func (p *projectRepositoryImpliment) SelectAll() ([]*model.Project, error) {
	queries := []qm.QueryMod{}
	projects, err := model.Projects(queries...).All(context.Background(), p.db)
	if err != nil {
		return nil, err
	}

	return projects, nil
}
