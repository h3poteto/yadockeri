package controllers

import (
	"net/http"
	"os"

	"github.com/gorilla/sessions"
	"github.com/h3poteto/yadockeri/app/usecases/user"
	"github.com/labstack/echo"
	"github.com/labstack/echo-contrib/session"
	"github.com/sirupsen/logrus"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/github"
)

type Oauth struct {
}

var githubConfig = oauth2.Config{
	ClientID:     os.Getenv("CLIENT_ID"),
	ClientSecret: os.Getenv("CLIENT_SECRET"),
	Scopes:       []string{"repo", "user:email"},
	Endpoint:     github.Endpoint,
}

func (o *Oauth) Login(c echo.Context) error {
	authURL := githubConfig.AuthCodeURL("state", oauth2.AccessTypeOffline)

	return c.Render(http.StatusOK, "login.html", map[string]interface{}{
		"title":    "yadockeri",
		"auth_url": authURL,
	})
}

func (o *Oauth) Callback(c echo.Context) error {
	code := c.QueryParam("code")
	token, err := githubConfig.Exchange(oauth2.NoContext, code)
	if err != nil {
		logrus.Error(err)
		return err
	}

	logrus.Debugf("token: %s", token)
	id, err := user.FindOrCreateUser(token.AccessToken)
	if err != nil {
		logrus.Error(err)
		return err
	}

	sess, _ := session.Get("login", c)
	sess.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   3600,
		HttpOnly: true,
	}
	sess.Values["current_user_id"] = id
	sess.Save(c.Request(), c.Response())
	return c.Redirect(http.StatusFound, "/")
}
