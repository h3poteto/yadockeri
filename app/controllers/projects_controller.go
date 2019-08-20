package controllers

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/asaskevich/govalidator"
	"github.com/h3poteto/yadockeri/app/middlewares"
	"github.com/h3poteto/yadockeri/app/usecases/project"
	"github.com/labstack/echo"
	"github.com/sirupsen/logrus"
)

type Projects struct{}

type NewProjectForm struct {
	Title             string                   `json:"title" form:"title" valid:"required,stringlength(1|255)"`
	BaseURL           string                   `json:"base_url" form:"base_url" valid:"required,stringlength(1|255)"`
	RepositoryOwner   string                   `json:"repository_owner" form:"repository_owner" valid:"required,stringlength(1|255)"`
	RepositoryName    string                   `json:"repository_name" form:"repository_name" valid:"required,stringlength(1|255)"`
	HelmRepositoryURL string                   `json:"helm_repository_url" form:"helm_repository_url" valid:"required,stringlength(1|255)"`
	HelmDirectoryName string                   `json:"helm_directory_name" form:"helm_directory_name" valid:"stringlength(0|255)"`
	Namespace         string                   `json:"namespace" form:"namespace" valid:"required,stringlength(1|255)"`
	ValueOptions      []*project.OverrideValue `json:"value_options" form:"value_options" valid:"-"`
}

type EditProjectForm struct {
	BaseURL           string                   `json:"base_url" form:"base_url" valid:"required,stringlength(1|255)"`
	HelmDirectoryName string                   `json:"helm_directory_name" form:"helm_directory_name" valid:"stringlength(0|255)"`
	Namespace         string                   `json:"namespace" form:"namespace" valid:"required,stringlength(1|255)"`
	ValueOptions      []*project.OverrideValue `json:"value_options" form:"value_options" valid:"-"`
}

func (p *Projects) Index(c echo.Context) error {
	projects, err := project.GetProjects()
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
	// Validate
	valid, err := govalidator.ValidateStruct(newProjectForm)
	logrus.Infof("Validation result: %v", valid)
	if err != nil {
		return err
	}

	proj, err := project.CreateProject(userID, newProjectForm.Title, newProjectForm.BaseURL, newProjectForm.RepositoryOwner, newProjectForm.RepositoryName, newProjectForm.HelmRepositoryURL, newProjectForm.HelmDirectoryName, newProjectForm.Namespace, newProjectForm.ValueOptions)
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

func (p *Projects) Update(c echo.Context) error {
	projectID, err := strconv.Atoi(c.Param("project_id"))
	if err != nil {
		return err
	}

	editProjectForm := new(EditProjectForm)
	err = c.Bind(editProjectForm)
	if err != nil {
		return err
	}
	// Validate
	valid, err := govalidator.ValidateStruct(editProjectForm)
	logrus.Infof("Validation result: %v", valid)
	if err != nil {
		return err
	}

	proj, err := project.UpdateProject(projectID, editProjectForm.BaseURL, editProjectForm.HelmDirectoryName, editProjectForm.Namespace, editProjectForm.ValueOptions)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, proj)
}
