package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func validatedBind(c echo.Context, i interface{}) error {
	err := c.Bind(i)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	err = c.Validate(i)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	return nil
}
