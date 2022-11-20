package domain

import (
	"testing"
	
	"github.com/stretchr/testify/assert"
)

func TestTask(t *testing.T) {
	task := &Task{ID: 1, Title: "title", Body: "body"}
	assert.Equal(t, 1, task.ID)
	assert.Equal(t, "title", task.Title)
	assert.Equal(t, "body", task.Body)
}

func TestTaskSet(t *testing.T) {
	task := &Task{ID: 1, Title: "title", Body: "body"}
	task.Set(task)
	assert.Equal(t, 1, task.ID)
	assert.Equal(t, "title", task.Title)
	assert.Equal(t, "body", task.Body)
}

func TestTaskSetFail(t *testing.T) {
	task := &Task{ID: 1, Title: "", Body: "body"}
	err := task.Set(task)
	assert.Error(t, err)
}