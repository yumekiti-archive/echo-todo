package repository

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/yumekiti/echo-todo/domain"
	mocks "github.com/yumekiti/echo-todo/mocks/domain/repository"
)

func TestTaskRepository(t *testing.T) {
	// Mock
	mockTaskRepo := new(mocks.TaskRepository)
	mockTask := &domain.Task{ID: 1, Title: "title", Body: "body"}
	mockTaskRepo.On("Get", 1).Return(mockTask, nil)
	mockTaskRepo.On("GetAll").Return([]*domain.Task{mockTask}, nil)
	mockTaskRepo.On("Save", mockTask).Return(mockTask, nil)
	mockTaskRepo.On("Remove", mockTask).Return(nil)
	mockTaskRepo.On("Update", mockTask).Return(mockTask, nil)

	// Test
	task, err := mockTaskRepo.Get(1)
	assert.NoError(t, err)
	assert.Equal(t, mockTask, task)

	tasks, err := mockTaskRepo.GetAll()
	assert.NoError(t, err)
	assert.Equal(t, []*domain.Task{mockTask}, tasks)

	task, err = mockTaskRepo.Save(mockTask)
	assert.NoError(t, err)
	assert.Equal(t, mockTask, task)

	err = mockTaskRepo.Remove(mockTask)
	assert.NoError(t, err)

	task, err = mockTaskRepo.Update(mockTask)
	assert.NoError(t, err)
	assert.Equal(t, mockTask, task)
}