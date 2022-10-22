package infra

import (
	"github.com/Rozelin-dc/SystemDesignTodoList/domain/model"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type TaskInfra struct {
	DB *sqlx.DB
}

func NewTaskInfra(db *sqlx.DB) *TaskInfra {
	return &TaskInfra{DB: db}
}

func (ti *TaskInfra) GetAllTasksByCreatorId(creatorId string) ([]*model.TaskSimple, error) {
	tasks := []*model.TaskSimple{}
	err := ti.DB.Select(
		&tasks,
		"SELECT `task_id`, `task_name`, `status` FROM `tasks` WHERE `creator_id` = ?",
		creatorId,
	)
	if err != nil {
		return nil, err
	}

	return tasks, nil
}

func (ti *TaskInfra) CreateTask(creatorId string, task *model.NewTask) (*model.TaskSimple, error) {
	uu, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}

	uuidStr := uu.String()

	if task.TimeLimit != nil {
		_, err = ti.DB.Exec(
			"INSERT INTO `tasks` (`task_id`, `creator_id`, `task_name`, `time_limit`) VALUES (?, ?, ?, ?)",
			uuidStr,
			creatorId,
			task.TaskName,
			task.TimeLimit.Local(),
		)
		if err != nil {
			return nil, err
		}
	} else {
		_, err = ti.DB.Exec(
			"INSERT INTO `tasks` (`task_id`, `creator_id`, `task_name`) VALUES (?, ?, ?)",
			uuidStr,
			creatorId,
			task.TaskName,
		)
		if err != nil {
			return nil, err
		}
	}

	return &model.TaskSimple{
		TaskId:    uuidStr,
		TaskName:  task.TaskName,
		Status:    0,
		TimeLimit: task.TimeLimit,
	}, nil
}

func (ti *TaskInfra) EditTask(task *model.TaskSimple) (*model.TaskSimple, error) {
	if task.TimeLimit != nil {
		_, err := ti.DB.Exec(
			"UPDATE `tasks` SET `task_name` = ?, `status` = ?, `time_limit` = ? WHERE `task_id` = ?",
			task.TaskName,
			task.Status,
			task.TimeLimit.Local(),
			task.TaskId,
		)
		if err != nil {
			return nil, err
		}
	} else {
		_, err := ti.DB.Exec(
			"UPDATE `tasks` SET `task_name` = ?, `status` = ? WHERE `task_id` = ?",
			task.TaskName,
			task.Status,
			task.TaskId,
		)
		if err != nil {
			return nil, err
		}
	}

	return task, nil
}

func (ti *TaskInfra) DeleteTask(taskId string, userId string) error {
	_, err := ti.DB.Exec(
		"DELETE FROM `tasks` WHERE `task_id` = ? AND `creator_id` = ?",
		taskId,
		userId,
	)

	return err
}
