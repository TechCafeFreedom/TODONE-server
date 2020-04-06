package project

import (
	"todone/db/mysql/model"
)

type Repository interface {
	InsertProject(project *model.Project) error
	SelectByPK(id int) (*model.Project, error)
	SelectByUserID(userID string) (model.ProjectSlice, error)
	SelectAll() (model.ProjectSlice, error)
}
