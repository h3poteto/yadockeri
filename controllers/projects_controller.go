package controllers

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	"github.com/lapras-inc/yadockeri/middlewares"
	"github.com/lapras-inc/yadockeri/usecases/project"
	"github.com/lapras-inc/yadockeri/values"
)

type Projects struct{}

type NewProjectForm struct {
	Title             string                  `json:"title" form:"title"`
	BaseURL           string                  `json:"base_url" form:"base_url"`
	RepositoryOwner   string                  `json:"repository_owner" form:"repository_owner"`
	RepositoryName    string                  `json:"repository_name" form:"repository_name"`
	HelmRepositoryUrl string                  `json:"helm_repository_url" form:"helm_repository_url"`
	HelmDirectoryName string                  `json:"helm_directory_name" form:"helm_directory_name"`
	Namespace         string                  `json:"namespace" form:"namespace"`
	ValueOptions      []*values.OverrideValue `json:"value_options" form:"value_options"`
}

func (p *Projects) Index(c echo.Context) error {
	uc, ok := c.(*middlewares.LoginContext)
	if !ok {
		return errors.New("Can not cast context")
	}
	user := uc.CurrentUser
	projects, err := project.GetProjectByUser(user.ID)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, projects)
}

func (p *Projects) Create(c echo.Context) error {
	uc, ok := c.(*middlewares.LoginContext)
	if !ok {
		return errors.New("Can not cast context")
	}
	userID := uc.CurrentUser.ID

	newProjectForm := new(NewProjectForm)
	err := c.Bind(newProjectForm)
	if err != nil {
		return err
	}
	proj, err := project.CreateProject(userID, newProjectForm.Title, newProjectForm.BaseURL, newProjectForm.RepositoryOwner, newProjectForm.RepositoryName, newProjectForm.HelmRepositoryUrl, newProjectForm.HelmDirectoryName, newProjectForm.Namespace, newProjectForm.ValueOptions)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusCreated, proj)
}

func (p *Projects) Show(c echo.Context) error {
	projectID, err := strconv.Atoi(c.Param("project_id"))
	if err != nil {
		return err
	}

	proj, err := project.GetProjectByID(projectID)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, proj)
}
