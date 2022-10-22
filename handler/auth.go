package handler

import (
	"net/http"

	"github.com/Rozelin-dc/SystemDesignTodoList/domain/model"
	"github.com/labstack/echo/v4"
)

func (h *Handler) PostLogin(c echo.Context) error {
	newUser := &model.UserSimple{}
	err := validatedBind(c, newUser)
	if err != nil {
		return err
	}

	user, err := h.ui.CheckRightUser(newUser)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	sess, err := h.si.CreateSession(user.UserId)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	c.SetCookie(&http.Cookie{
		Name:  "session_id",
		Value: sess.SessionId,
	})

	return c.NoContent(http.StatusOK)
}
