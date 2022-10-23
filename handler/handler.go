package handler

import (
	"net/http"

	"github.com/Rozelin-dc/SystemDesignTodoList/domain/repository"
	"github.com/Rozelin-dc/SystemDesignTodoList/infra"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
)

type Handler struct {
	ui repository.UserRepository
	ti repository.TaskRepository
	si repository.SessionRepository
}

func NewHandler(db *sqlx.DB) *Handler {
	ui := infra.NewUserInfra(db)
	ti := infra.NewTaskInfra(db)
	si := infra.NewSessionInfra(db)
	return &Handler{
		ui: ui,
		ti: ti,
		si: si,
	}
}

func (h *Handler) NotImpl(c echo.Context) error {
	return c.NoContent(http.StatusNotImplemented)
}
