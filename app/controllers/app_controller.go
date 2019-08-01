package controllers

import (
	"net/http"

	"github.com/labstack/echo"
)

type App struct{}

func (a *App) Index(c echo.Context) error {
	return c.Render(http.StatusOK, "index.html", map[string]interface{}{})
}
