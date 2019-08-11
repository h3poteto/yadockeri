package controllers

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/asaskevich/govalidator"
	"github.com/h3poteto/yadockeri/app/middlewares"
	"github.com/h3poteto/yadockeri/app/usecases/branch"
	"github.com/labstack/echo"
	"github.com/sirupsen/logrus"
)

type Branches struct{}

type NewBranchForm struct {
	Name string `json:"name" form:"name" valid:"required,stringlength(1|255)"`
}

type StatusResponse struct {
	Status string `json:"status"`
}

func (b *Branches) Index(c echo.Context) error {
	projectID, err := strconv.Atoi(c.Param("project_id"))
	if err != nil {
		return err
	}
	branches, err := branch.GetBranchesByProject(projectID)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, branches)
}

func (b *Branches) Create(c echo.Context) error {
	uc, ok := c.(*middlewares.LoginContext)
	if !ok {
		return errors.New("Can not cast context")
	}
	user := uc.CurrentUser
	projectID, err := strconv.Atoi(c.Param("project_id"))
	if err != nil {
		return err
	}

	newBranchForm := new(NewBranchForm)
	err = c.Bind(newBranchForm)
	if err != nil {
		return err
	}
	// Validate
	valid, err := govalidator.ValidateStruct(newBranchForm)
	logrus.Infof("Validation result: %v", valid)
	if err != nil {
		return err
	}

	branch, err := branch.CreateBranch(projectID, user.ID, newBranchForm.Name)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusCreated, branch)
}

func (b *Branches) Show(c echo.Context) error {
	projectID, err := strconv.Atoi(c.Param("project_id"))
	if err != nil {
		return err
	}
	branchID, err := strconv.Atoi(c.Param("branch_id"))
	if err != nil {
		return err
	}
	findBranch, err := branch.GetBranch(projectID, branchID)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, findBranch)
}

func (b *Branches) Deploy(c echo.Context) error {
	projectID, err := strconv.Atoi(c.Param("project_id"))
	if err != nil {
		return err
	}
	branchID, err := strconv.Atoi(c.Param("branch_id"))
	if err != nil {
		return err
	}

	status, err := branch.Deploy(projectID, branchID)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, &StatusResponse{
		Status: status,
	})
}

func (b *Branches) Status(c echo.Context) error {
	projectID, err := strconv.Atoi(c.Param("project_id"))
	if err != nil {
		return err
	}
	branchID, err := strconv.Atoi(c.Param("branch_id"))
	if err != nil {
		return err
	}

	status, err := branch.GetStatus(projectID, branchID)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, &StatusResponse{
		Status: status,
	})
}

func (b *Branches) Delete(c echo.Context) error {
	projectID, err := strconv.Atoi(c.Param("project_id"))
	if err != nil {
		return err
	}
	branchID, err := strconv.Atoi(c.Param("branch_id"))
	if err != nil {
		return err
	}

	err = branch.Delete(projectID, branchID)

	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, "")
}
