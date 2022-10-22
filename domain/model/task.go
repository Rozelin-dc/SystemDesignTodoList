package model

import "time"

type TaskSimple struct {
	TaskId   string `json:"taskId" db:"task_id"`
	TaskName string `json:"taskName" db:"task_name"`
	Status   int    `json:"status" db:"status"`
}

type Task struct {
	TaskSimple

	CreatorId string    `db:"creator_id"`
	CreatedAt time.Time `db:"created_at"`
}
