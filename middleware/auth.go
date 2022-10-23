package middleware

import (
	"net/http"

	"github.com/Rozelin-dc/SystemDesignTodoList/domain/model"
	"github.com/Rozelin-dc/SystemDesignTodoList/handler"
	"github.com/labstack/echo/v4"
)

func EnsureAuthorized(h *handler.Handler) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			sess := &model.Session{}
			err := h.PickSession(c, sess)
			if err != nil {
				return err
			}

			return next(c)
		}
	}
}

func EnsureAccessRightToAccount(h *handler.Handler) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			sess := &model.Session{}
			err := h.PickSession(c, sess)
			if err != nil {
				return err
			}

			uid := c.Param("uid")
			if uid == "" {
				return echo.NewHTTPError(http.StatusBadRequest, "`uid` is required")
			}

			if uid != sess.UserId {
				return echo.NewHTTPError(http.StatusForbidden, "cannot access other's account")
			}

			return next(c)
		}
	}
}