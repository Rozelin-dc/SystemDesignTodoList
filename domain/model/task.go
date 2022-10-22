package model

import "time"

type TaskSimple struct {
	TaskId    string     `json:"taskId" db:"task_id"`
	TaskName  string     `json:"taskName" db:"task_name"`
	Status    int        `json:"status" db:"status"`
	TimeLimit *time.Time `json:"timeLimit" db:"time_limit"`
}

type NewTask struct {
	TaskName  string     `json:"taskName" db:"task_name"`
	TimeLimit *time.Time `json:"timeLimit" db:"time_limit"`
}

type Task struct {
	TaskSimple

	CreatorId string    `db:"creator_id"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}
