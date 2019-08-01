package controllers

import (
	"errors"
	"net/http"

	"github.com/labstack/echo"
	"github.com/lapras-inc/yadockeri/middlewares"
	"github.com/lapras-inc/yadockeri/usecases/github"
)

type Github struct{}

func (g *Github) Repos(c echo.Context) error {
	uc, ok := c.(*middlewares.LoginContext)
	if !ok {
		return errors.New("Can not cast context")
	}
	token := uc.CurrentUser.OauthToken
	repos, err := github.GetRepos(token)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, repos)
}

func (g *Github) Branches(c echo.Context) error {
	uc, ok := c.(*middlewares.LoginContext)
	if !ok {
		return errors.New("Can not cast context")
	}
	token := uc.CurrentUser.OauthToken

	owner := c.QueryParam("owner")
	repo := c.QueryParam("repo")
	branches, err := github.GetBranches(owner, repo, token)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, branches)
}
