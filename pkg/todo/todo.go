package todo

import "time"

type TodoTask interface {
	ToDTO() TodoDTO
}

type todoState struct {
	Id          string
	Title       string
	Description string
	CreatedAt   time.Time
	Status      TodoStatus
}

type TodoToCreate struct {
	Id          string `json:"id" binding:"required"`
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
}

func New(data TodoToCreate) TodoTask {
	return &todoState{
		Id:          data.Id,
		Title:       data.Title,
		Description: data.Description,
		Status:      Todo,
		CreatedAt:   time.Now(),
	}
}
