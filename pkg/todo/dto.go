package todo

import "time"

type TodoDTO struct {
	Id          string     `json:"id" binding:"required"`
	Title       string     `json:"title" binding:"required"`
	Description string     `json:"description" binding:"required"`
	CreatedAt   string     `json:"createdAt" binding:"required"`
	Status      TodoStatus `json:"status" binding:"required"`
}

func FromDTO(data TodoDTO) (TodoTask, error) {
	createdAt, err := time.Parse("2006-01-02", data.CreatedAt)

	if err != nil {
		return nil, err
	}

	return &todoState{
		Id:          data.Id,
		Title:       data.Title,
		Description: data.Description,
		Status:      data.Status,
		CreatedAt:   createdAt,
	}, nil
}

func (state *todoState) ToDTO() TodoDTO {
	var TodoDTO = TodoDTO{
		Title:       state.Title,
		Description: state.Description,
		Id:          state.Id,
		Status:      state.Status,
		CreatedAt:   state.CreatedAt.Format("2006-01-02"),
	}

	return TodoDTO
}
