package services

import (
	"gorotines/models"
	"gorotines/pkg/repositories"
	"log"
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
func (s *TaskService) CreateTask(task *models.Task, resultChan chan<- error) {
	err := s.repository.Create(task)
	resultChan <- err
}

// GetTasks return all tasks
func (s *TaskService) GetTasks(resultChan chan<- []models.Task) {
	tasks, err := s.repository.GetAll()
	if err != nil {
		log.Println("Failed to fetch tasks: ", err)
		resultChan <- nil
		return
	}
	resultChan <- tasks
}

// GetTaskByID return one task by ID
func (s *TaskService) GetTaskByID(id int, resultChan chan<- *models.Task) {
	task, err := s.repository.GetByID(id)
	if err != nil {
		log.Println("Failed to fetch task: ", err)
		resultChan <- nil
		return
	}
	resultChan <- task
}

// UpdateTask a task by ID
func (s *TaskService) UpdateTask(task *models.Task, resultChan chan<- error) {
	err := s.repository.Update(task)
	resultChan <- err
}

// DeleteTask excludes a task by ID
func (s *TaskService) DeleteTask(id int, resultChan chan<- error) {
	err := s.repository.Delete(id)
	resultChan <- err
}
