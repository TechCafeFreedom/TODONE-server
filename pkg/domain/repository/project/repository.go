package project

import (
	model "todone/db/mysql/models"
)

type Repository interface {
	InsertProject(project *model.Project) error
	SelectAll() ([]*model.Project, error)
}
