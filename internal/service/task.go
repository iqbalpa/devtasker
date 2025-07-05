package service

import (
	"devtasker/internal/model"
	"devtasker/internal/repository"
	"fmt"
)

type ITaskService interface {
	CreateTask(title, description string) (*model.Task, error)
	GetTaskByID(id string) (*model.Task, error)
	GetAllTasks() ([]*model.Task, error)
	UpdateTask(id, title, description string, status model.TaskStatus) (*model.Task, error)
	DeleteTask(id string) (*model.Task, error)
}

type TaskService struct {
	r repository.ITaskRepository
}

func New(r *repository.ITaskRepository) *TaskService {
	return &TaskService{
		r: *r,
	}
}

func (ts *TaskService) CreateTask(title, description string) (*model.Task, error) {
	if title == "" || description == "" {
		return nil, fmt.Errorf("title and description cannot be empty")
	}
	t, err := ts.r.CreateTask(title, description)
	if err != nil {
		return nil, err
	}
	return t, nil
}

func (ts *TaskService) GetTaskByID(id string) (*model.Task, error) {
	t, err := ts.r.GetTaskByID(id)
	if err != nil {
		return nil, err
	}
	return t, nil
}

func (ts *TaskService) GetAllTasks() ([]*model.Task, error) {
	tasks, err := ts.r.GetAllTasks()
	if err != nil {
		return []*model.Task{}, nil
	}
	return tasks, nil
}

func (ts *TaskService) UpdateTask(id, title, description string, status model.TaskStatus) (*model.Task, error) {
	if title == "" || description == "" || status == "" {
		return nil, fmt.Errorf("title and description cannot be empty")
	}
	t, err := ts.r.UpdateTask(id, title, description, status)
	if err != nil {
		return nil, err
	}
	return t, nil
}

func (ts *TaskService) DeleteTask(id string) (*model.Task, error) {
	t, err := ts.r.DeleteTask(id)
	if err != nil {
		return nil, err
	}
	return t, nil
}
