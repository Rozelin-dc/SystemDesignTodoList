package handler

import (
	"fmt"
	"net/http"

	"github.com/Rozelin-dc/SystemDesignTodoList/domain/model"
	"github.com/labstack/echo/v4"
)

func (h *Handler) PickTaskById(taskId string, task *model.Task) error {
	task, err := h.ti.GetTaskByTaskId(taskId)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	if task == nil {
		return echo.NewHTTPError(http.StatusNotFound, fmt.Sprintf("no such task `%s`", taskId))
	}
	return nil
}
