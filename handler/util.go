package handler

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

func validatedBind(c echo.Context, i interface{}) error {
	err := c.Bind(i) // リクエストボディの取り出し
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	err = c.Validate(i) // バリデーションの実施
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
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
		Expires:  time.Date(2030, 12, 31, 23, 59, 59, 99, time.Local),
		HttpOnly: true,
	})

	return nil
}
