package usecase

import (
	"github.com/yumekiti/echo-todo/domain"
	"github.com/yumekiti/echo-todo/domain/repository"
)

// TaskUsecase task usecaseのinterface
type TaskUsecase interface {
	Get(id int) (*domain.Task, error)
	GetAll() ([]*domain.Task, error)
	Save(*domain.Task) (*domain.Task, error)
	Remove(id int) error
	Update(id int, task *domain.Task) (*domain.Task, error)
}

type taskUsecase struct {
	taskRepo repository.TaskRepository
}

func NewTaskUsecase(taskRepo repository.TaskRepository) TaskUsecase {
	return &taskUsecase{taskRepo: taskRepo}
}

func (tu *taskUsecase) Get(id int) (*domain.Task, error) {
	foundTask, err := tu.taskRepo.Get(id)
	if err != nil {
		return nil, err
	}

	return foundTask, nil
}

func (tu *taskUsecase) GetAll() ([]*domain.Task, error) {
	foundTask, err := tu.taskRepo.GetAll()
	if err != nil {
		return nil, err
	}

	return foundTask, nil
}

func (tu *taskUsecase) Save(task *domain.Task) (*domain.Task, error) {
	savedTask, err := tu.taskRepo.Save(task)
	if err != nil {
		return nil, err
	}

	return savedTask, nil
}

func (tu *taskUsecase) Remove(id int) error {
	task, err := tu.taskRepo.Get(id)
	if err != nil {
		return err
	}

	err = tu.taskRepo.Remove(task)
	if err != nil {
		return err
	}

	return nil
}

func (tu *taskUsecase) Update(id int, task *domain.Task) (*domain.Task, error) {
	targetTask, err := tu.taskRepo.Get(id)
	if err != nil {
		return nil, err
	}

	updatedTask, err := tu.taskRepo.Update(targetTask)
	if err != nil {
		return nil, err
	}

	return updatedTask, nil
}