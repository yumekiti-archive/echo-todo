package repository

import (
	"github.com/yumekiti/echo-todo/domain"
)

type TaskRepository interface {
	Get(id int) (*domain.Task, error)
	GetAll() ([]*domain.Task, error)
	Save(*domain.Task) (*domain.Task, error)
	Remove(*domain.Task) error
	Update(*domain.Task) (*domain.Task, error)
}
