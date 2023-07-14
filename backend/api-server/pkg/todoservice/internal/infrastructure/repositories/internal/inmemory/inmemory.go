package inmemory

import "golang-practice/pkg/todoservice/internal/domain/entities/todo"

type repositoryStateInMemory struct {
	data []todo.TodoState
}

func New() todo.TodoRepository {
	return &repositoryStateInMemory{
		data: []todo.TodoState{},
	}
}

func (state *repositoryStateInMemory) Save(todoTask todo.TodoTask) error {
	state.data = append(state.data, todoTask.CopyState())
	return nil
}

func (state *repositoryStateInMemory) GetById(id string) (todo.TodoTask, error) {

	for _, TodoState := range state.data {
		if TodoState.Id == id {
			var todoTask = todo.FromState(TodoState)
			return todoTask, nil
		}
	}

	return nil, nil
}

func (state *repositoryStateInMemory) GetAll() ([]todo.TodoTask, error) {

	var todoList []todo.TodoTask

	for _, TodoState := range state.data {
		var todoTask = todo.FromState(TodoState)
		todoList = append(todoList, todoTask)
	}

	return todoList, nil
}
