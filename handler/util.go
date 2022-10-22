package handler

import (
	"net/http"

	"github.com/Rozelin-dc/SystemDesignTodoList/domain/model"
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

func pickSession(c echo.Context, h *Handler, sess *model.Session) error {
	cookie, err := c.Cookie("session_id")
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	sess, err = h.si.CheckSession(cookie.Value)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	if sess != nil {
		return onFailedError(c)
	}

	return nil
}
