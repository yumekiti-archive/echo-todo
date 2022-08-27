package infrastructure

import (
	"gorm.io/gorm"

	"github.com/yumekiti/echo-todo/domain"
	"github.com/yumekiti/echo-todo/domain/repository"
)

type taskRepository struct {
	Conn *gorm.DB
}

func NewTaskRepository(conn *gorm.DB) repository.TaskRepository {
	return &taskRepository{Conn: conn}
}

func (tr *taskRepository) GetAll() ([]*domain.Task, error) {
	var tasks []*domain.Task
	if err := tr.Conn.Find(&tasks).Error; err != nil {
		return nil, err
	}
	return tasks, nil
}

func (tr *taskRepository) Get(id int) (*domain.Task, error) {
	task := &domain.Task{ID: id}
	if err := tr.Conn.First(&task, id).Error; err != nil {
		return nil, err
	}
	return task, nil
}

func (tr *taskRepository) Save(task *domain.Task) (*domain.Task, error) {
	if err := tr.Conn.Save(&task).Error; err != nil {
		return nil, err
	}
	return task, nil
}

func (tr *taskRepository) Remove(task *domain.Task) error {
	if err := tr.Conn.Delete(&task).Error; err != nil {
		return err
	}
	return nil
}

func (tr *taskRepository) Update(task *domain.Task) (*domain.Task, error) {
	if err := tr.Conn.Model(&task).Save(&task).Error; err != nil {
		return nil, err
	}
	return task, nil
}
