package project

import (
	"errors"
	"net/http"
	"todone/db/mysql/model"
	"todone/pkg/api/request/reqbody"
	"todone/pkg/api/response"
	projectinteractor "todone/pkg/api/usecase/project"

	"github.com/gin-gonic/gin"
)

type Server struct {
	projectInteractor projectinteractor.Interactor
}

func New(projectInteractor projectinteractor.Interactor) Server {
	return Server{projectInteractor: projectInteractor}
}

func (s *Server) CreateNewProject(ctx *gin.Context) {
	var reqBody reqbody.ProjectCreate
	if err := ctx.BindJSON(&reqBody); err != nil {
		ctx.Error(err)
	}

	userID, ok := ctx.Get("AUTHED_USER_ID")
	if !ok {
		ctx.Error(errors.New("userID is not found in context"))
	}

	if err := s.projectInteractor.CreateNewProject(ctx, userID.(string), reqBody.Title, reqBody.Description); err != nil {
		ctx.Error(err)
	}

	ctx.Status(http.StatusOK)
}

func (s *Server) GetProjectByPK(ctx *gin.Context) {
	var reqBody reqbody.ProjectPK
	if err := ctx.BindJSON(&reqBody); err != nil {
		ctx.Error(err)
	}

	project, err := s.projectInteractor.GetByPK(ctx, reqBody.ID)
	if err != nil {
		ctx.Error(err)
	}

	ctx.JSON(http.StatusOK, convertToProjectResponse(project))
}

func (s *Server) GetProjectsByUserID(ctx *gin.Context) {
	userID, ok := ctx.Get("AUTHED_USER_ID")
	if !ok {
		ctx.Error(errors.New("userID is not found in context"))
	}

	projects, err := s.projectInteractor.GetByUserID(ctx, userID.(string))
	if err != nil {
		ctx.Error(err)
	}

	ctx.JSON(http.StatusOK, convertToProjectsResponse(projects))
}

func (s *Server) GetAllProjects(ctx *gin.Context) {
	projects, err := s.projectInteractor.GetAll(ctx)
	if err != nil {
		ctx.Error(err)
	}

	ctx.JSON(http.StatusOK, convertToProjectsResponse(projects))
}

func convertToProjectResponse(project *model.Project) response.ProjectResponse {
	return response.ProjectResponse{
		ID:          project.ID,
		Title:       project.Title,
		Description: project.Description,
		CreatedAt:   project.CreatedAt,
		UpdatedAt:   project.UpdatedAt,
	}
}

func convertToProjectsResponse(projects model.ProjectSlice) response.ProjectsResponse {
	res := make(response.ProjectsResponse, 0, len(projects))
	for _, data := range projects {
		res = append(res, &response.ProjectResponse{
			ID:          data.ID,
			Title:       data.Title,
			Description: data.Description,
			CreatedAt:   data.CreatedAt,
			UpdatedAt:   data.UpdatedAt,
		})
	}
	return res
}
