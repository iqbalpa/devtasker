package model

type TaskStatus string

const (
	Pending    TaskStatus = "pending"
	InProgress TaskStatus = "in-progress"
	Completed  TaskStatus = "completed"
	Archived   TaskStatus = "archived"
	Deleted    TaskStatus = "deleted"
)

type Task struct {
	ID          string     `json:"id"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	Status      TaskStatus `json:"status"`
	CreatedAt   string     `json:"created_at"`
}
