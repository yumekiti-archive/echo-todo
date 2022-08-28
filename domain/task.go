package domain

import "errors"

type Task struct {
	ID    int
	Title string
	Body  string
}

func (t *Task) Set(task *Task) error {
	// Validation
	if task.Title == "" {
		return errors.New("Please enter a title")
	}

	t.Title = task.Title
	t.Body = task.Body

	return nil
}
