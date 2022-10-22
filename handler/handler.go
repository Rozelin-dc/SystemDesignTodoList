package handler

import (
	"net/http"

	"github.com/Rozelin-dc/SystemDesignTodoList/domain/model"
	"github.com/Rozelin-dc/SystemDesignTodoList/infra"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
)

type Handler struct {
	UserInfra    *infra.UserInfra
	TaskInfra    *infra.TaskInfra
	SessionInfra *infra.SessionInfra
}

func NewHandler(db *sqlx.DB) *Handler {
	ui := infra.NewUserInfra(db)
	ti := infra.NewTaskInfra(db)
	si := infra.NewSessionInfra(db)
	return &Handler{
		UserInfra:    ui,
		TaskInfra:    ti,
		SessionInfra: si,
	}
}

func (h *Handler) NotImpl(c echo.Context) error {
	return c.NoContent(http.StatusOK)
}

func (h *Handler) PostLogin(c echo.Context) error {
	newUser := &model.UserSimple{}
	err := validatedBind(c, newUser)
	if err != nil {
		return err
	}

	user, err := h.UserInfra.CheckRightUser(newUser)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	sess, err := h.SessionInfra.CreateSession(user.UserId)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	c.SetCookie(&http.Cookie{
		Name:  "session_id",
		Value: sess.SessionId,
	})

	return c.NoContent(http.StatusOK)
}
