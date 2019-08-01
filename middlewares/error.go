package middlewares

import (
	"github.com/labstack/echo"
	"github.com/sirupsen/logrus"
)

func ErrorLogging(e *echo.Echo) func(error, echo.Context) {
	return func(err error, c echo.Context) {
		logrus.Error(err)
		e.DefaultHTTPErrorHandler(err, c)
	}
}
