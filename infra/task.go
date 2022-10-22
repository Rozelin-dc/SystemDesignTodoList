package infra

import (
	"github.com/Rozelin-dc/SystemDesignTodoList/domain/model"
	"github.com/Rozelin-dc/SystemDesignTodoList/domain/repository"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type taskInfra struct {
	db *sqlx.DB
}

func NewTaskInfra(db *sqlx.DB) repository.TaskRepository {
	return &taskInfra{db: db}
}

func (ti *taskInfra) GetAllTasksByCreatorId(creatorId string) ([]*model.TaskSimple, error) {
	tasks := []*model.TaskSimple{}
	err := ti.db.Select(
		&tasks,
		"SELECT `task_id`, `task_name`, `status` FROM `tasks` WHERE `creator_id` = ?",
		creatorId,
	)
	if err != nil {
		return nil, err
	}

	return tasks, nil
}

func (ti *taskInfra) CreateTask(creatorId string, task *model.NewTask) (*model.TaskSimple, error) {
	uu, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}

	uuidStr := uu.String()

	if task.TimeLimit != nil {
		_, err = ti.db.Exec(
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
		_, err = ti.db.Exec(
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

func (ti *taskInfra) EditTask(task *model.TaskSimple) (*model.TaskSimple, error) {
	if task.TimeLimit != nil {
		_, err := ti.db.Exec(
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
		_, err := ti.db.Exec(
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

func (ti *taskInfra) DeleteTask(taskId string, userId string) error {
	_, err := ti.db.Exec(
		"DELETE FROM `tasks` WHERE `task_id` = ? AND `creator_id` = ?",
		taskId,
		userId,
	)

	return err
}
