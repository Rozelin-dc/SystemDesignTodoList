package repository

import "github.com/Rozelin-dc/SystemDesignTodoList/domain/model"

type TaskRepository interface {
	GetAllTasksByCreatorId(creatorId string) ([]*model.TaskSimple, error)
	CreateTask(creatorId string, taskName string) (*model.TaskSimple, error)
	EditTask(task *model.TaskSimple) (*model.TaskSimple, error)
	DeleteTask(taskId string) error
}