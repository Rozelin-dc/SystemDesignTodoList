package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func bindParm(c echo.Context, i interface{}) error {
	err := c.Bind(i)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return nil
}

func createSessionAndSetCookie(c echo.Context, h *Handler, userId string) error {
	sess, err := h.si.CreateSession(userId)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	c.SetCookie(&http.Cookie{
		Name:     "session_id",
		Value:    sess.SessionId,
		Path:     "/",
		HttpOnly: true,
	})

	return nil
}
