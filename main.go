package main

import (
	"io"
	"os"
	"text/template"

	"github.com/gorilla/sessions"
	"github.com/h3poteto/yadockeri/app/controllers"
	"github.com/h3poteto/yadockeri/app/middlewares"
	"github.com/labstack/echo"
	"github.com/labstack/echo-contrib/session"
)

//go:generate go-assets-builder --output=config/bindata.go -s="/config" -p=config config/settings.yml

// TemplateRenderer is a custom html/template renderer for Echo framework
type TemplateRenderer struct {
	templates *template.Template
}

// Render renders a template document
func (t *TemplateRenderer) Render(w io.Writer, name string, data interface{}, c echo.Context) error {

	// Add global methods if data is a map
	if viewContext, isMap := data.(map[string]interface{}); isMap {
		viewContext["reverse"] = c.Echo().Reverse
	}

	return t.templates.ExecuteTemplate(w, name, data)
}

func main() {
	e := echo.New()
	e.Use(session.Middleware(sessions.NewCookieStore([]byte(os.Getenv("SESSION_SECRET")))))
	e.Use(middlewares.CustomizeLogger())

	renderer := &TemplateRenderer{
		templates: template.Must(template.ParseGlob("app/templates/*.html")),
	}
	e.Renderer = renderer
	e.HTTPErrorHandler = middlewares.ErrorLogging(e)

	e.Static("/assets", "assets")
	oauth := controllers.Oauth{}
	e.GET("/login", oauth.Login)
	e.GET("/oauth/callback", oauth.Callback)

	login := e.Group("/")
	login.Use(middlewares.Login())

	projects := controllers.Projects{}
	login.GET("api/v1/projects", projects.Index)
	login.POST("api/v1/projects", projects.Create)
	login.GET("api/v1/projects/:project_id", projects.Show)
	login.PATCH("api/v1/projects/:project_id", projects.Update)

	github := controllers.Github{}
	login.GET("api/v1/github/repos", github.Repos)
	login.GET("api/v1/github/branches", github.Branches)

	branches := controllers.Branches{}
	login.GET("api/v1/projects/:project_id/branches", branches.Index)
	login.GET("api/v1/projects/:project_id/branches/:branch_id", branches.Show)
	login.POST("api/v1/projects/:project_id/branches", branches.Create)
	login.PATCH("api/v1/projects/:project_id/branches/:branch_id/deploy", branches.Deploy)
	login.GET("api/v1/projects/:project_id/branches/:branch_id/status", branches.Status)
	login.DELETE("api/v1/projects/:project_id/branches/:branch_id", branches.Delete)

	app := controllers.App{}
	login.GET("", app.Index)
	login.GET("*", app.Index)

	e.Logger.Fatal(e.Start(":9090"))
}
