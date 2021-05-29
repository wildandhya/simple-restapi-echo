package helper

import (
	"fmt"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo"
)

func ErrHandler(e error, c echo.Context) {
	report, ok := e.(*echo.HTTPError)
	if !ok {
		report = echo.NewHTTPError(http.StatusInternalServerError, map[string]string{"message": e.Error()})
	}

	if customError, ok := e.(validator.ValidationErrors); ok {
		for _, err := range customError {
			switch err.Tag() {
			case "required":
				report.Message = fmt.Sprintf("%s is required", err.Field())
			case "email":
				report.Message = fmt.Sprintf("%s is not valid email", err.Field())
			case "gte":
				report.Message = fmt.Sprintf("%s must be greater then %s", err.Field(), err.Param())
			case "lte":
				report.Message = fmt.Sprintf("%s must be lower then %s", err.Field(), err.Param())
			}
		}
	}
	c.Logger().Error(report)
	c.JSON(report.Code, report)
}
