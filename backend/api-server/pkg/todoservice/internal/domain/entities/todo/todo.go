package todo

import "time"

type TodoTask interface {
	CopyState() TodoState
}

type TodoState struct {
	Id          string
	Title       string
	Description string
	CreatedBy   string
	CreatedAt   time.Time
	Status      TodoStatus
}

type TodoToCreate struct {
	Id          string
	Title       string
	Description string
	CreatedBy   string
}

func New(data TodoToCreate) TodoTask {
	return &TodoState{
		Id:          data.Id,
		Title:       data.Title,
		Description: data.Description,
		Status:      Todo,
		CreatedBy:   data.CreatedBy,
		CreatedAt:   time.Now(),
	}
}

func (state *TodoState) CopyState() TodoState {
	var copiedState = TodoState{
		Title:       state.Title,
		Description: state.Description,
		Id:          state.Id,
		Status:      state.Status,
		CreatedAt:   state.CreatedAt,
		CreatedBy:   state.CreatedBy,
	}

	return copiedState
}

func FromState(data TodoState) TodoTask {
	return &TodoState{
		Id:          data.Id,
		Title:       data.Title,
		Description: data.Description,
		Status:      data.Status,
		CreatedAt:   data.CreatedAt,
		CreatedBy:   data.CreatedBy,
	}
}
