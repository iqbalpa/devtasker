package model

import "time"

type TaskStatus string

const (
	Pending    TaskStatus = "pending"
	InProgress TaskStatus = "in-progress"
	Completed  TaskStatus = "completed"
	Archived   TaskStatus = "archived"
	Deleted    TaskStatus = "deleted"
)

type Task struct {
	ID          string     `json:"id" gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	Status      TaskStatus `json:"status"`
	CreatedAt   time.Time  `json:"created_at"`
	UpdatedAt   time.Time  `json:"updated_at"`
}
