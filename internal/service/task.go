package service

import (
	"context"
	"devtasker/internal/model"
	"devtasker/internal/repository"
	"devtasker/internal/utils"
	"fmt"
)

type ITaskService interface {
	GetAllTasks(ctx context.Context) ([]model.Task, error)
	GetTaskByID(ctx context.Context, id string) (model.Task, error)
	CreateTask(ctx context.Context, title, description string) (model.Task, error)
	UpdateTask(ctx context.Context, id, title, description string, status model.TaskStatus) (model.Task, error)
	DeleteTask(ctx context.Context, id string) (model.Task, error)
}

type TaskService struct {
	r repository.ITaskRepository
}

func New(r repository.ITaskRepository) *TaskService {
	return &TaskService{
		r: r,
	}
}

func (ts *TaskService) GetAllTasks(ctx context.Context) ([]model.Task, error) {
	username, _ := ctx.Value(utils.UsernameKey).(string)
	tasks, err := ts.r.GetAllTasks(username)
	if err != nil {
		return []model.Task{}, nil
	}
	return tasks, nil
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

func (ts *TaskService) UpdateTask(ctx context.Context, id, title, description string, status model.TaskStatus) (model.Task, error) {
	if title == "" || description == "" || status == "" {
		return model.Task{}, fmt.Errorf("title and description cannot be empty")
	}
	// Step 1: Get task by ID
	task, err := ts.r.GetTaskByID(id)
	if err != nil {
		return model.Task{}, fmt.Errorf("task not found")
	}
	// Step 2: Check ownership
	username, _ := ctx.Value(utils.UsernameKey).(string)
	if task.UserUsername != username {
		return model.Task{}, fmt.Errorf("you don't have permission to access this task")
	}
	// Step 3: Update the task
	t, err := ts.r.UpdateTask(id, title, description, status)
	if err != nil {
		return model.Task{}, err
	}
	return t, nil
}

func (ts *TaskService) DeleteTask(ctx context.Context, id string) (model.Task, error) {
	// Step 1: Get task by ID
	task, err := ts.r.GetTaskByID(id)
	if err != nil {
		return model.Task{}, fmt.Errorf("task not found")
	}
	// Step 2: Check ownership
	username, _ := ctx.Value(utils.UsernameKey).(string)
	if task.UserUsername != username {
		return model.Task{}, fmt.Errorf("you don't have permission to access this task")
	}
	// Step 3: Delete the task
	t, err := ts.r.DeleteTask(id)
	if err != nil {
		return model.Task{}, err
	}
	return t, nil
}
