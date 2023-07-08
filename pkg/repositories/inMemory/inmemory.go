package inmemory

import (
	"golang-practice/pkg/todo"
)

type RepositoryStateInMemory struct {
	data []todo.TodoDTO
}

func New() todo.TodoRepository {
	return &RepositoryStateInMemory{
		data: []todo.TodoDTO{},
	}
}

func (state *RepositoryStateInMemory) Save(todoTask todo.TodoTask) error {
	state.data = append(state.data, todoTask.ToDTO())
	return nil
}

func (state *RepositoryStateInMemory) GetById(id string) (todo.TodoTask, error) {

	for _, todoDTO := range state.data {
		if todoDTO.Id == id {

			var todoTask, err = todo.FromDTO(todoDTO)

			if err != nil {
				return nil, err
			}

			return todoTask, nil
		}
	}

	return nil, nil
}

func (state *RepositoryStateInMemory) GetAll() ([]todo.TodoTask, error) {

	var todoList []todo.TodoTask

	for _, todoDTO := range state.data {
		var todoTask, err = todo.FromDTO(todoDTO)

		if err != nil {
			return nil, err
		}

		todoList = append(todoList, todoTask)
	}

	return todoList, nil
}
