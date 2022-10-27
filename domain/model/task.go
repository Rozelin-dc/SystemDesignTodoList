package model

import "time"

type TaskSimple struct {
	TaskId    string `json:"taskId" db:"task_id"`
	TaskName  string `json:"taskName" db:"task_name" validate:"taskName"`
	Status    int    `json:"status" db:"status" validate:"status"`
	TimeLimit string `json:"timeLimit" db:"time_limit" validate:"omitempty,timeLimit"`
}

type TaskList struct {
	HasNext bool           `json:"hasNext"`
	Tasks   *[]*TaskSimple `json:"tasks"`
}

type TaskUpdate struct {
	TaskName  string `json:"taskName" db:"task_name" validate:"taskName"`
	Status    int    `json:"status" db:"status" validate:"status"`
	TimeLimit string `json:"timeLimit" db:"time_limit" validate:"omitempty,timeLimit"`
}

type NewTask struct {
	TaskName  string `json:"taskName" db:"task_name" validate:"taskName"`
	TimeLimit string `json:"timeLimit" db:"time_limit" validate:"omitempty,timeLimit"`
}

type Task struct {
	TaskSimple

	CreatorId string    `db:"creator_id"`
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}
