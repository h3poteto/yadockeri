package middlewares

import (
	"fmt"
	"os"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// CustomizeLogger prepqre my logger for echo
func CustomizeLogger() echo.MiddlewareFunc {
	return middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: printColored("status") + "=${status} " + printColored("method") + "=${method} " + printColored("path") + "=${uri} " + printColored("requestID") + "=${id} " + printColored("latency") + "=${latency_human} " + printColored("time") + "=${time_rfc3339_nano}\n",
		Output: os.Stdout,
	})
}

func printColored(str string) string {
	return fmt.Sprintf("\x1b[%dm%s\x1b[0m", 34, str)
}
