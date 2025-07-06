package repository

import (
	"devtasker/internal/model"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ITaskRepository interface {
	CreateTask(title, description string) (model.Task, error)
	GetTaskByID(id string) (model.Task, error)
	GetAllTasks() ([]model.Task, error)
	UpdateTask(id, title, description string, status model.TaskStatus) (model.Task, error)
	DeleteTask(id string) (model.Task, error)
}

type TaskRepository struct {
	db *gorm.DB
}

func New(db *gorm.DB) *TaskRepository {
	return &TaskRepository{
		db: db,
	}
}

func (tr *TaskRepository) CreateTask(title, description string) (model.Task, error) {
	id := uuid.NewString()
	t := model.Task{
		ID:          id,
		Title:       title,
		Description: description,
		Status:      model.Pending,
		CreatedAt:   time.Now().String(),
	}
	tr.db.Save(t)
	return t, nil
}

func (tr *TaskRepository) GetTaskByID(id string) (model.Task, error) {
	var task model.Task
	result := tr.db.First(&task, "id = ?", id)
	if result.Error != nil {
		return model.Task{}, result.Error
	}
	return task, nil
}

func (tr *TaskRepository) GetAllTasks() ([]model.Task, error) {
	var tasks []model.Task
	result := tr.db.Find(&tasks)
	if result.Error != nil {
		return []model.Task{}, result.Error
	}
	return tasks, nil
}

func (tr *TaskRepository) UpdateTask(id, title, description string, status model.TaskStatus) (model.Task, error) {
	result := tr.db.Model(&model.Task{}).Where("id = ?", id).Updates(map[string]interface{}{
		"title":       title,
		"description": description,
		"status":      status,
	})
	if result.Error != nil {
		return model.Task{}, result.Error
	}
	return model.Task{}, nil
}

func (tr *TaskRepository) DeleteTask(id string) (model.Task, error) {
	var task model.Task
	result := tr.db.Where("id = ?", id).Delete(task)
	if result.Error != nil {
		return model.Task{}, result.Error
	}
	return task, nil
}
