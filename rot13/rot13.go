package rot13

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/masc23/go-algorithms-echo/core"
)

func rot13(c echo.Context) error {
	requestResponseObject := new(core.RequestResponseObject)

	if err := c.Bind(requestResponseObject); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	requestResponseObject.Output = rot13String(requestResponseObject.Input)

	return c.JSON(http.StatusOK, requestResponseObject)
}

func init() {
	core.EchoInstance.POST("/rot13", rot13)
}
