package usecase

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/yumekiti/echo-todo/domain"
	mocks "github.com/yumekiti/echo-todo/mocks/domain/repository"
)

func TestGet(t *testing.T) {
	// Mock
	mockTaskRepo := new(mocks.TaskRepository)
	mockTask := &domain.Task{ID: 1, Title: "title", Body: "body"}
	mockTaskRepo.On("Get", mock.AnythingOfType("int")).Return(mockTask, nil)

	// UseCase
	taskUsecase := NewTaskUsecase(mockTaskRepo)

	// Test
	task, err := taskUsecase.Get(1)
	assert.NoError(t, err)
	assert.Equal(t, mockTask, task)
}

func TestGetAll(t *testing.T) {
	// Mock
	mockTaskRepo := new(mocks.TaskRepository)
	mockTasks := []*domain.Task{
		{ID: 1, Title: "title", Body: "body"},
		{ID: 2, Title: "title", Body: "body"},
	}
	mockTaskRepo.On("GetAll").Return(mockTasks, nil)

	// UseCase
	taskUsecase := NewTaskUsecase(mockTaskRepo)

	// Test
	tasks, err := taskUsecase.GetAll()
	assert.NoError(t, err)
	assert.Equal(t, mockTasks, tasks)
}

func TestSave(t *testing.T) {
	// Mock
	mockTaskRepo := new(mocks.TaskRepository)
	mockTask := &domain.Task{ID: 1, Title: "title", Body: "body"}
	mockTaskRepo.On("Save", mock.AnythingOfType("*domain.Task")).Return(mockTask, nil)

	// UseCase
	taskUsecase := NewTaskUsecase(mockTaskRepo)

	// Test
	task, err := taskUsecase.Save(mockTask)
	assert.NoError(t, err)
	assert.Equal(t, mockTask, task)
}

func TestSaveMulti(t *testing.T) {
	// Mock
	mockTaskRepo := new(mocks.TaskRepository)
	mockTasks := []*domain.Task{
		{ID: 1, Title: "title", Body: "body"},
		{ID: 2, Title: "title", Body: "body"},
	}
	for _, task := range mockTasks {
		mockTaskRepo.On("Save", task).Return(task, nil)
	}

	// UseCase
	taskUsecase := NewTaskUsecase(mockTaskRepo)

	// Test
	for _, task := range mockTasks {
		task, err := taskUsecase.Save(task)
		assert.NoError(t, err)
		assert.Equal(t, task, task)
	}
}

func TestRemove(t *testing.T) {
	// Mock
	mockTaskRepo := new(mocks.TaskRepository)
	mockTask := &domain.Task{ID: 1, Title: "title", Body: "body"}
	mockTaskRepo.On("Get", mock.AnythingOfType("int")).Return(mockTask, nil)
	mockTaskRepo.On("Save", mock.AnythingOfType("*domain.Task")).Return(mockTask, nil)
	mockTaskRepo.On("Remove", mock.AnythingOfType("*domain.Task")).Return(nil)

	// UseCase
	taskUsecase := NewTaskUsecase(mockTaskRepo)

	// Test
	err := taskUsecase.Remove(1)
	assert.NoError(t, err)
}

func TestUpdate(t *testing.T) {
	// Mock
	mockTaskRepo := new(mocks.TaskRepository)
	mockTask := &domain.Task{ID: 1, Title: "title", Body: "body"}
	mockTaskRepo.On("Get", mock.AnythingOfType("int")).Return(mockTask, nil)
	mockTaskRepo.On("Save", mock.AnythingOfType("*domain.Task")).Return(mockTask, nil)
	mockTaskRepo.On("Update", mock.AnythingOfType("*domain.Task")).Return(mockTask, nil)

	// UseCase
	taskUsecase := NewTaskUsecase(mockTaskRepo)

	// Test
	task, err := taskUsecase.Update(mockTask)
	assert.NoError(t, err)
	assert.Equal(t, mockTask, task)
}