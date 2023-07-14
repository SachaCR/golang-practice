package todoservice

import (
	"golang-practice/pkg/todoservice/internal/domain/entities/todo"
	"time"
)

type TodoDTO struct {
	Id          string `json:"id" binding:"required"`
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
	CreatedAt   string `json:"createdAt" binding:"required"`
	Status      string `json:"status" binding:"required"`
}

type CreateTodoDTO struct {
	Id          string `json:"id" binding:"required"`
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
}

func FromDTO(data TodoDTO) (todo.TodoTask, error) {
	createdAt, err := time.Parse("2006-01-02", data.CreatedAt)

	if err != nil {
		return nil, err
	}

	var todoTask = todo.FromState(todo.TodoState{
		Id:          data.Id,
		Title:       data.Title,
		Description: data.Description,
		Status:      todo.TodoStatusFromString(data.Status),
		CreatedAt:   createdAt,
	})

	return todoTask, nil
}

func ToDTO(todoTask todo.TodoTask) TodoDTO {
	var todoState = todoTask.CopyState()
	var todoDTO = TodoDTO{
		Id:          todoState.Id,
		Title:       todoState.Title,
		Description: todoState.Description,
		Status:      todoState.Status.String(),
		CreatedAt:   todoState.CreatedAt.Format("2006-01-02"),
	}

	return todoDTO
}
