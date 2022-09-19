package base64

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/masc23/go-algorithms-echo/core"
)

func encode(c echo.Context) error {
	requestResponseObject := new(core.RequestResponseObject)

	if err := c.Bind(requestResponseObject); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	requestResponseObject.Output = encodeString(requestResponseObject.Input)

	return c.JSON(http.StatusOK, requestResponseObject)
}

func decode(c echo.Context) error {
	requestResponseObject := new(core.RequestResponseObject)

	if err := c.Bind(requestResponseObject); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	requestResponseObject.Output = decodeString(requestResponseObject.Input)

	return c.JSON(http.StatusOK, requestResponseObject)
}

func init() {
	core.EchoInstance.POST("/base64/encode", encode)
	core.EchoInstance.POST("/base64/decode", decode)
}
