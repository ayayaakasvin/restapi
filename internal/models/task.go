package models

import "time"

type Task struct {
	ID          int64     `json:"taskId"`
	UserID      int64     `json:"userId"`
	TaskContent string    `json:"taskContent"`
	CreatedAt   time.Time `json:"createdAt"`
}