package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/yumekiti/echo-todo/domain"
	"github.com/yumekiti/echo-todo/usecase"
)

type TaskHandler interface {
	Post() echo.HandlerFunc
	Get() echo.HandlerFunc
	GetAll() echo.HandlerFunc
	Put() echo.HandlerFunc
	Delete() echo.HandlerFunc
}

type taskHandler struct {
	taskUsecase usecase.TaskUsecase
}

func NewTaskHandler(taskUsecase usecase.TaskUsecase) TaskHandler {
	return &taskHandler{taskUsecase: taskUsecase}
}

type requestTask struct {
	Title string `json:"title"`
	Body  string `json:"body"`
}

type responseTask struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Body  string `json:"body"`
}

func (th *taskHandler) Post() echo.HandlerFunc {
	return func(c echo.Context) error {
		var req requestTask
		if err := c.Bind(&req); err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		createdTask, err := th.taskUsecase.Save(&domain.Task{ID: 0, Title: req.Title, Body: req.Body})
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		res := responseTask{
			Title: createdTask.Title,
			Body:  createdTask.Body,
		}

		return c.JSON(http.StatusCreated, res)
	}
}

func (th *taskHandler) Get() echo.HandlerFunc {
	return func(c echo.Context) error {
		id, err := strconv.Atoi((c.Param("id")))
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		foundTask, err := th.taskUsecase.Get(id)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		res := responseTask{
			ID:    foundTask.ID,
			Title: foundTask.Title,
			Body:  foundTask.Body,
		}

		return c.JSON(http.StatusOK, res)
	}
}

func (th *taskHandler) GetAll() echo.HandlerFunc {
	return func(c echo.Context) error {
		foundTask, err := th.taskUsecase.GetAll()
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		res := []responseTask{}
		for _, task := range foundTask {
			res = append(res, responseTask{
				ID:    task.ID,
				Title: task.Title,
				Body:  task.Body,
			})
		}

		return c.JSON(http.StatusOK, res)
	}
}

func (th *taskHandler) Put() echo.HandlerFunc {
	return func(c echo.Context) error {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		var req requestTask
		if err := c.Bind(&req); err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		fmt.Print(req.Title)

		updatedTask, err := th.taskUsecase.Update(&domain.Task{ID: id, Title: req.Title, Body: req.Body})
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		res := responseTask{
			ID:    updatedTask.ID,
			Title: updatedTask.Title,
			Body:  updatedTask.Body,
		}

		return c.JSON(http.StatusOK, res)
	}
}

func (th *taskHandler) Delete() echo.HandlerFunc {
	return func(c echo.Context) error {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		err = th.taskUsecase.Remove(id)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		return c.NoContent(http.StatusNoContent)
	}
}
