package todoservice

import (
	"golang-practice/pkg/todoservice/internal/domain/entities/todo"
	"golang-practice/pkg/todoservice/internal/infrastructure/repositories"

	"github.com/spf13/viper"
)

type TodoService interface {
	AddTodo(createTodoDTO CreateTodoDTO) (TodoDTO, error)
	GetTodoById(id string) (TodoDTO, error)
	GetTodoList() ([]TodoDTO, error)
	//updateTodo(id string, data TodoUpdate)
}

type todoServiceState struct {
	TodoRepository todo.TodoRepository
}

var emptyDTO = TodoDTO{}

func New() TodoService {

	var repositoryEngineString = viper.GetString("todoService.repositories.engine")
	var repositoryEngine = repositories.RepositoryEngineFromString(repositoryEngineString)
	var todoRepository = repositories.New(repositoryEngine)

	return &todoServiceState{
		TodoRepository: todoRepository,
	}
}

func (state *todoServiceState) AddTodo(createTodoDTO CreateTodoDTO) (TodoDTO, error) {

	var todoTask = todo.New(todo.TodoToCreate{
		Id:          createTodoDTO.Id,
		Title:       createTodoDTO.Title,
		Description: createTodoDTO.Description,
	})

	var error = state.TodoRepository.Save(todoTask)

	if error != nil {
		return emptyDTO, error
	}

	return ToDTO(todoTask), nil
}

func (state *todoServiceState) GetTodoById(id string) (TodoDTO, error) {

	var todoTask, error = state.TodoRepository.GetById(id)

	if error != nil {
		return emptyDTO, error
	}

	if todoTask == nil {
		return emptyDTO, TodoNotFound()
	}

	return ToDTO(todoTask), nil
}

func (state *todoServiceState) GetTodoList() ([]TodoDTO, error) {

	var todoTaskList, error = state.TodoRepository.GetAll()

	if error != nil {
		return []TodoDTO{}, error
	}

	var todoDTOs = []TodoDTO{}

	for _, todoTask := range todoTaskList {
		todoDTOs = append(todoDTOs, ToDTO(todoTask))
	}

	return todoDTOs, nil
}
