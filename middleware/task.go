package middleware

import (
	"net/http"

	"github.com/Rozelin-dc/SystemDesignTodoList/domain/model"
	"github.com/Rozelin-dc/SystemDesignTodoList/handler"
	"github.com/labstack/echo/v4"
)

func EnsureExistTaskAndHaveAccessRight(h *handler.Handler) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			sess, err := h.PickSession(c)
			if err != nil {
				return err
			}

			tid := c.Param("tid")
			if tid == "" {
				return echo.NewHTTPError(http.StatusBadRequest, "`tid` is required")
			}

			task := &model.Task{}
			err = h.PickTaskById(tid, task)
			if err != nil {
				return err
			}

			if sess.UserId != task.CreatorId {
				return echo.NewHTTPError(http.StatusForbidden, "cannot access other's task")
			}

			return next(c)
		}
	}
}
