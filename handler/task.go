package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/Rozelin-dc/SystemDesignTodoList/domain/model"
	"github.com/labstack/echo/v4"
)

func (h *Handler) PostTask(c echo.Context) error {
	sess, err := h.PickSession(c)
	if err != nil {
		return err
	}

	task := &model.NewTask{}
	err = bindParm(c, task)
	if err != nil {
		return err
	}

	createdTask, err := h.ti.CreateTask(sess.UserId, task)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, createdTask)
}

func (h *Handler) GetTasks(c echo.Context) error {
	sess, err := h.PickSession(c)
	if err != nil {
		return err
	}

	l := c.QueryParam("limit")
	if l == "" {
		l = "20"
	}
	limit, err := strconv.Atoi(l)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	o := c.QueryParam("offset")
	if o == "" {
		o = "0"
	}
	offset, err := strconv.Atoi(o)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	tasks, err := h.ti.GetAllTasksByCreatorId(sess.UserId, limit, offset)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, tasks)
}

func (h *Handler) PatchTask(c echo.Context) error {
	tid := c.Param("tid")
	task := &model.TaskUpdate{}
	err := bindParm(c, task)
	if err != nil {
		return err
	}

	edited, err := h.ti.EditTask(tid, task)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, edited)
}

func (h *Handler) DeleteTask(c echo.Context) error {
	tid := c.Param("tid")
	sess, err := h.PickSession(c)
	if err != nil {
		return err
	}

	err = h.ti.DeleteTaskByTaskId(tid, sess.UserId)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.NoContent(http.StatusOK)
}

func (h *Handler) PickTaskById(taskId string) (*model.Task, error) {
	task, err := h.ti.GetTaskByTaskId(taskId)
	if err != nil {
		return nil, echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	if task == nil {
		return nil, echo.NewHTTPError(http.StatusNotFound, fmt.Sprintf("no such task `%s`", taskId))
	}
	return task, nil
}
