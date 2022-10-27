package repository

import "github.com/Rozelin-dc/SystemDesignTodoList/domain/model"

type TaskRepository interface {
	GetAllTasksByCreatorId(creatorId string, limit int, offset int) (*model.TaskList, error)
	GetTaskByTaskId(taskId string) (*model.Task, error)
	CreateTask(creatorId string, task *model.NewTask) (*model.TaskSimple, error)
	EditTask(taskId string, task *model.TaskUpdate) (*model.TaskSimple, error)
	DeleteTaskByTaskId(taskId string, userId string) error
	DeleteTaskByUserId(userId string) error
}
