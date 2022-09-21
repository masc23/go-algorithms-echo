package mathematics

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/masc23/go-algorithms-echo/core"
)

func mathHandler(c echo.Context) error {
	requestResponseObject := new(core.RequestResponseObject)

	if err := c.Bind(requestResponseObject); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	switch action := c.Param("action"); action {
	case "evaluate-equation":
		requestResponseObject.Output = evaluateEquation(requestResponseObject.Input)

	case "regula-falsi":
		requestResponseObject.Output = runRegulaFalsi(requestResponseObject.Input)

	default:
		return echo.NewHTTPError(http.StatusNotFound, "action not found")
	}

	return c.JSON(http.StatusOK, requestResponseObject)
}

func init() {
	core.EchoInstance.POST("/math/:action", mathHandler)
}
