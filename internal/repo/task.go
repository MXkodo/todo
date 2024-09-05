package repo

import (
	"github.com/MXkodo/todo/model"
	"gorm.io/gorm"
)

type TaskRepo interface {
	Create(task *model.Task) (*model.Task, error)
	FindAll() ([]model.Task, error)
	FindByID(id uint) (*model.Task, error)
	Update(task *model.Task) (*model.Task, error)
	Delete(id uint) error
}

type taskRepo struct {
	db *gorm.DB
}

func NewTaskRepo(db *gorm.DB) TaskRepo {
	return &taskRepo{db: db}
}

func (r *taskRepo) Create(task *model.Task) (*model.Task, error) {
	result := r.db.Create(task)
	return task, result.Error
}

func (r *taskRepo) FindAll() ([]model.Task, error) {
	var tasks []model.Task
	result := r.db.Find(&tasks)
	return tasks, result.Error
}

func (r *taskRepo) FindByID(id uint) (*model.Task, error) {
	var task model.Task
	result := r.db.First(&task, id)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, nil 
		}
		return nil, result.Error
	}
	return &task, nil
}

func (r *taskRepo) Update(task *model.Task) (*model.Task, error) {
	result := r.db.Save(task)
	return task, result.Error
}

func (r *taskRepo) Delete(id uint) error {
	result := r.db.Delete(&model.Task{}, id)
	return result.Error
}
