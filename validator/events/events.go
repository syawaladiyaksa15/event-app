package events

import (
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

var ErroHandlerEvent = func(err error, c echo.Context) {

	report, ok := err.(*echo.HTTPError)
	if !ok {
		report = echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	if castedObject, ok := err.(validator.ValidationErrors); ok {
		for _, err := range castedObject {
			switch err.Tag() {
			case "required":
				report.Message = fmt.Sprintf("%s is required",
					err.Field())
			case "min":
				report.Message = fmt.Sprintf("%s Minimum format %s characters",
					err.Field(), err.Param())
			case "number":
				report.Message = fmt.Sprintf("%s value required number",
					err.Field())
			}

			break
		}
	}

	c.Logger().Error(report)
	c.JSON(report.Code, report)
}
