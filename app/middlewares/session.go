package middlewares

import (
	"errors"
	"net/http"
	"strings"

	"github.com/h3poteto/yadockeri/app/usecases/user"
	"github.com/labstack/echo"
	"github.com/labstack/echo-contrib/session"
)

type LoginContext struct {
	echo.Context
	CurrentUser *user.User
}

func Login() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			user, err := checkLogin(c)
			if err != nil {
				// requestがjsonを求めているならjsonを返す
				// そうでないならリダイレクトする
				if isJSONRequest(c) {
					c.JSON(http.StatusUnauthorized, err)
				} else {
					c.Redirect(http.StatusFound, "/login")
				}
				return err
			}
			uc := &LoginContext{
				c,
				user,
			}
			return next(uc)
		}
	}
}

func checkLogin(c echo.Context) (*user.User, error) {
	sess, _ := session.Get("login", c)
	if sess.Values["current_user_id"] == nil {
		return nil, errors.New("session not found")
	}
	return user.LoginConfirm(sess.Values["current_user_id"].(int))
}

func isJSONRequest(c echo.Context) bool {
	headers := c.Request().Header
	if len(headers[echo.HeaderContentType]) > 0 {
		for _, h := range headers[echo.HeaderContentType] {
			if strings.Contains(h, echo.MIMEApplicationJSON) {
				return true
			}
		}
	}
	return false
}
