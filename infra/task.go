package infra

import (
	"database/sql"
	"errors"
	"fmt"

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

func (ti *taskInfra) GetAllTasksByCreatorId(creatorId string, limit int, offset int, status string, name string) (*model.TaskList, error) {
	query := "SELECT `task_id`, `task_name`, `status`, `time_limit` FROM `tasks` WHERE `creator_id` = ?"

	if status != "" { // 状態による検索条件が指定されていた場合
		query = fmt.Sprintf("%s AND `status` = %s", query, status)
	}

	if name != "" { // 名前による検索条件が指定されていた場合
		query = fmt.Sprintf("%s AND `task_name` LIKE '%%%s%%'", query, name)
	}

	query = fmt.Sprintf("%s ORDER BY `created_at` DESC LIMIT ? OFFSET ?", query)

	tasks := []*model.TaskSimple{}
	err := ti.db.Select(
		&tasks,
		query,
		creatorId,
		limit,
		offset,
	)
	if err != nil {
		return nil, err
	}
	if len(tasks) == 0 {
		return &model.TaskList{
			HasNext: false,
			Tasks:   nil,
		}, nil
	}

	count := 0
	err = ti.db.Get(
		&count,
		"SELECT COUNT(*) FROM `tasks` WHERE `creator_id` = ? GROUP BY `creator_id`",
		creatorId,
	)
	if err != nil {
		return nil, err
	}

	return &model.TaskList{
		HasNext: count > len(tasks)+offset,
		Tasks:   &tasks,
	}, nil
}

func (ti *taskInfra) GetTaskByTaskId(taskId string) (*model.Task, error) {
	task := model.Task{}
	err := ti.db.Get(
		&task,
		"SELECT * FROM `tasks` WHERE `task_id` = ?",
		taskId,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}

		return nil, err
	}

	return &task, nil
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
			task.TimeLimit,
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

func (ti *taskInfra) EditTask(taskId string, task *model.TaskUpdate) (*model.TaskSimple, error) {
	if task.TimeLimit != nil {
		_, err := ti.db.Exec(
			"UPDATE `tasks` SET `task_name` = ?, `status` = ?, `time_limit` = ? WHERE `task_id` = ?",
			task.TaskName,
			task.Status,
			task.TimeLimit,
			taskId,
		)
		if err != nil {
			return nil, err
		}
	} else {
		_, err := ti.db.Exec(
			"UPDATE `tasks` SET `task_name` = ?, `status` = ? WHERE `task_id` = ?",
			task.TaskName,
			task.Status,
			taskId,
		)
		if err != nil {
			return nil, err
		}
	}

	return &model.TaskSimple{
		TaskId:    taskId,
		TaskName:  task.TaskName,
		Status:    task.Status,
		TimeLimit: task.TimeLimit,
	}, nil
}

func (ti *taskInfra) DeleteTaskByTaskId(taskId string, userId string) error {
	_, err := ti.db.Exec(
		"DELETE FROM `tasks` WHERE `task_id` = ? AND `creator_id` = ?",
		taskId,
		userId,
	)

	return err
}

func (ti *taskInfra) DeleteTaskByUserId(userId string) error {
	_, err := ti.db.Exec(
		"DELETE FROM `tasks` WHERE `creator_id` = ?",
		userId,
	)

	return err
}
