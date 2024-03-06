package repositories

import (
	"gorotines/db"
	"gorotines/models"
)

type TaskRepository struct{}

// represents a tasks repository
func NewTaskRepository() *TaskRepository {
	return &TaskRepository{}
}

// NewTaskRepository create a nre instancy of TaskRepository
func (r *TaskRepository) Create(task *models.Task) error {
	return db.DB.Create(task).Error
}

// GetAll returns all tasks
func (r *TaskRepository) GetAll() ([]models.Task, error) {
	var tasks []models.Task
	err := db.DB.Find(&tasks).Error
	if err != nil {
		return nil, err
	}
	return tasks, nil
}

// GetByID returns one task by ID
func (r *TaskRepository) GetByID(id int) (*models.Task, error) {
	var task models.Task
	err := db.DB.First(&task, id).Error
	if err != nil {
		return nil, err
	}
	return &task, nil
}

func (r *TaskRepository) Update(task *models.Task) error {
	return db.DB.Save(task).Error
}

// Delete deletes a task
func (r *TaskRepository) Delete(id int) error {
	return db.DB.Delete(&models.Task{}, id).Error
}
