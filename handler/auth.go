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
	if user == nil {
		return echo.NewHTTPError(http.StatusNotFound, "no such user")
	}

	err = createSessionAndSetCookie(c, h, user.UserId)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, user)
}

func (h *Handler) PostLogout(c echo.Context) error {
	cookie, err := c.Cookie("session_id")
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	err = h.si.DeleteSessionBySessionId(cookie.Value)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	return c.NoContent(http.StatusOK)
}

func (h *Handler) PickSession(c echo.Context) (*model.Session, error) {
	cookie, err := c.Cookie("session_id")
	if err != nil {
		return nil, echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	sess, err := h.si.CheckSession(cookie.Value)
	if err != nil {
		return nil, echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	if sess == nil {
		return nil, echo.NewHTTPError(http.StatusUnauthorized, "unauthorized")
	}

	return sess, nil
}
