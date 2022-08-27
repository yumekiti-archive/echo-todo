package domain

import "errors"

type Task struct {
	ID    int
	Title string
	Body  string
}

func (t *Task) Validation(task *Task) error {
	if task.Title == "" {
		return errors.New("Please enter a title")
	}

	t.Title = task.Title
	t.Body = task.Body

	return nil
}
