package service

import (
	"github.com/MXkodo/todo/internal/repo"
	"github.com/MXkodo/todo/model"
)

type TaskService interface {
	Create(task *model.Task) (*model.Task, error)
	FindAll() ([]model.Task, error)
	FindByID(id uint) (*model.Task, error)
	Update(id uint, task *model.Task) (*model.Task, error)
	Delete(id uint) error
}

type taskService struct {
	taskRepo repo.TaskRepo
}

func NewTaskService(taskRepo repo.TaskRepo) TaskService {
	return &taskService{taskRepo: taskRepo}
}

func (s *taskService) Create(task *model.Task) (*model.Task, error) {
	return s.taskRepo.Create(task)
}

func (s *taskService) FindAll() ([]model.Task, error) {
	return s.taskRepo.FindAll()
}

func (s *taskService) FindByID(id uint) (*model.Task, error) {
	return s.taskRepo.FindByID(id)
}

func (s *taskService) Update(id uint, task *model.Task) (*model.Task, error) {
	existingTask, err := s.taskRepo.FindByID(id)
	if err != nil {
		return nil, err
	}
	existingTask.Title = task.Title
	existingTask.Description = task.Description
	existingTask.DueDate = task.DueDate
	return s.taskRepo.Update(existingTask)
}

func (s *taskService) Delete(id uint) error {
	return s.taskRepo.Delete(id)
}
