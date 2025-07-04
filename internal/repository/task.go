package repository

import (
	"devtasker/internal/model"
	"fmt"
	"time"

	"github.com/google/uuid"
)

type ITaskRepository interface {
	CreateTask(title, description string) (*model.Task, error)
	GetTaskByID(id string) (*model.Task, error)
	GetAllTasks() ([]*model.Task, error)
	UpdateTask(id, title, description string, status model.TaskStatus) (*model.Task, error)
	DeleteTask(id string) (*model.Task, error)
	// ArchiveTask(id string) error
	// UnarchiveTask(id string) error
	// MarkTaskAsCompleted(id string) error
	// MarkTaskAsInProgress(id string) error
}

type TaskRepository struct {
	Tasks map[string]*model.Task
}

func New() *TaskRepository {
	return &TaskRepository{
		Tasks: make(map[string]*model.Task),
	}
}

func (tr *TaskRepository) CreateTask(title, description string) (*model.Task, error) {
	id := uuid.NewString()
	t := model.Task{
		ID:          id,
		Title:       title,
		Description: description,
		Status:      model.Pending,
		CreatedAt:   time.Now().String(),
	}
	tr.Tasks[id] = &t
	return &t, nil
}

func (tr *TaskRepository) GetTaskByID(id string) (*model.Task, error) {
	t, ok := tr.Tasks[id]
	if !ok {
		return nil, fmt.Errorf("task with id %s is not found", id)
	}
	return t, nil
}

func (tr *TaskRepository) GetAllTasks() ([]*model.Task, error) {
	tasks := []*model.Task{}
	for _, v := range tr.Tasks {
		tasks = append(tasks, v)
	}
	return tasks, nil
}

func (tr *TaskRepository) UpdateTask(id, title, description string, status model.TaskStatus) (*model.Task, error) {
	t := tr.Tasks[id]
	t.Title = title
	t.Description = description
	t.Status = status
	return t, nil
}

func (tr *TaskRepository) DeleteTask(id string) (*model.Task, error) {
	t, ok := tr.Tasks[id]
	if !ok {
		return nil, fmt.Errorf("task with id %s is not found", id)
	}
	delete(tr.Tasks, id)
	return t, nil
}
