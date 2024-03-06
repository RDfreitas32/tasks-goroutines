package services

import (
	"gorotines/models"
	"gorotines/pkg/repositories"
)

// TaskService represents a task service
type TaskService struct {
	repository repositories.TaskRepository
}

// NewTaskService creates a new intance od TaskService
func NewTaskService(repository repositories.TaskRepository) *TaskService {
	return &TaskService{repository}
}

// CreateTask creates a new task
func (s *TaskService) CreateTask(task *models.Task) error {
	return s.repository.Create(task)
}

// GetTasks return all tasks
func (s *TaskService) GetTasks() ([]models.Task, error) {
	return s.repository.GetAll()
}

// GetTaskByID return one task by ID
func (s *TaskService) GetTaskByID(id int) (*models.Task, error) {
	return s.repository.GetByID(id)
}

// UpdateTask a task by ID
func (s *TaskService) UpdateTask(task *models.Task) error {
	return s.repository.Update(task)
}

// DeleteTask excludes a task by ID
func (s *TaskService) DeleteTask(id int) error {
	return s.repository.Delete(id)
}
