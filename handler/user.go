package handler

import (
	"fmt"
	"net/http"

	"github.com/Rozelin-dc/SystemDesignTodoList/domain/model"
	"github.com/labstack/echo/v4"
)

func (h *Handler) PostUser(c echo.Context) error {
	newUser := &model.UserSimple{}
	err := validatedBind(c, newUser)
	if err != nil {
		return err
	}

	err = checkUserNameDuplicated(c, h, newUser.UserName)
	if err != nil {
		return err
	}

	createdUser, err := h.ui.CreateUser(newUser)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	createSessionAndSetCookie(c, h, createdUser.UserId)

	return c.JSON(http.StatusOK, createdUser)
}

func (h *Handler) GetUserMe(c echo.Context) error {
	sess := &model.Session{}
	err := h.PickSession(c, sess)
	if err != nil {
		return err
	}

	user, err := h.ui.GetUser(sess.UserId)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, user)
}

func (h *Handler) PatchUser(c echo.Context) error {
	uid := c.Param("uid")
	editUser := &model.UserUpdate{}
	err := validatedBind(c, editUser)
	if err != nil {
		return err
	}

	user, err := h.ui.EditUser(uid, editUser)
	if err != nil {
		echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, user)
}

func (h *Handler) DeleteUser(c echo.Context) error {
	uid := c.Param("uid")
	pass := c.QueryParam("password")
	if pass == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "`password` is required")
	}

	err := h.ui.DeleteUser(uid, pass)
	if err != nil {
		echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.NoContent(http.StatusOK)
}

func checkUserNameDuplicated(c echo.Context, h *Handler, userName string) error {
	user, err := h.ui.CheckUsedUserName(userName)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	if user != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("user name %s is duplicated", userName))
	}

	return nil
}
