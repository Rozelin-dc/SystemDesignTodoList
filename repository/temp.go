package repository

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (repo *Repository) NotImpl(c echo.Context) error {
	return c.NoContent(http.StatusOK)
}
