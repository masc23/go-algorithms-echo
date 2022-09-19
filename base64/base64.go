package base64

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/masc23/go-algorithms-echo/core"
)

func base64Handler(c echo.Context) error {
	requestResponseObject := new(core.RequestResponseObject)

	if err := c.Bind(requestResponseObject); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	switch action := c.Param("action"); action {
	case "decode":
		requestResponseObject.Output = decodeString(requestResponseObject.Input)

	case "encode":
		requestResponseObject.Output = encodeString(requestResponseObject.Input)

	default:
		return echo.NewHTTPError(http.StatusNotFound, "action not found")
	}

	return c.JSON(http.StatusOK, requestResponseObject)
}

func init() {
	core.EchoInstance.POST("/base64/:action", base64Handler)
}
