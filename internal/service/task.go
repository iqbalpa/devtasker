package service

import (
	"context"
	"devtasker/internal/model"
	"devtasker/internal/repository"
	"devtasker/internal/utils"
	"fmt"
)

type ITaskService interface {
	CreateTask(ctx context.Context, title, description string) (model.Task, error)
	GetTaskByID(ctx context.Context, id string) (model.Task, error)
	GetAllTasks(ctx context.Context) ([]model.Task, error)
	UpdateTask(id, title, description string, status model.TaskStatus) (model.Task, error)
	DeleteTask(id string) (model.Task, error)
}

type TaskService struct {
	r repository.ITaskRepository
}

func New(r repository.ITaskRepository) *TaskService {
	return &TaskService{
		r: r,
	}
}

func (ts *TaskService) CreateTask(ctx context.Context, title, description string) (model.Task, error) {
	if title == "" || description == "" {
		return model.Task{}, fmt.Errorf("title and description cannot be empty")
	}
	username, _ := ctx.Value(utils.UsernameKey).(string)
	t, err := ts.r.CreateTask(username, title, description)
	if err != nil {
		return model.Task{}, err
	}
	return t, nil
}

func (ts *TaskService) GetTaskByID(ctx context.Context, id string) (model.Task, error) {
	t, err := ts.r.GetTaskByID(id)
	if err != nil {
		return model.Task{}, err
	}
	username, _ := ctx.Value(utils.UsernameKey).(string)
	if t.UserUsername != username {
		return model.Task{}, fmt.Errorf("you don't have permission to access this task")
	}
	return t, nil
}

func (ts *TaskService) GetAllTasks(ctx context.Context) ([]model.Task, error) {
	username, _ := ctx.Value(utils.UsernameKey).(string)
	tasks, err := ts.r.GetAllTasks(username)
	if err != nil {
		return []model.Task{}, nil
	}
	return tasks, nil
}

func (ts *TaskService) UpdateTask(id, title, description string, status model.TaskStatus) (model.Task, error) {
	if title == "" || description == "" || status == "" {
		return model.Task{}, fmt.Errorf("title and description cannot be empty")
	}
	t, err := ts.r.UpdateTask(id, title, description, status)
	if err != nil {
		return model.Task{}, err
	}
	return t, nil
}

func (ts *TaskService) DeleteTask(id string) (model.Task, error) {
	t, err := ts.r.DeleteTask(id)
	if err != nil {
		return model.Task{}, err
	}
	return t, nil
}
