package core

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var EchoInstance *echo.Echo = echo.New()

func Start() {
	EchoInstance.Use(middleware.Logger())
	EchoInstance.Use(middleware.Recover())
	EchoInstance.Logger.Fatal(EchoInstance.Start(":1323"))
}
