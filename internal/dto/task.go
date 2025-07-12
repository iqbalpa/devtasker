package dto

import "devtasker/internal/model"

type CreateTaskRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

type UpdateTaskRequest struct {
	Title       string           `json:"title"`
	Description string           `json:"description"`
	Status      model.TaskStatus `json:"status"`
}
