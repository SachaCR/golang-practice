package todoservice

import (
	"errors"
	"golang-practice/pkg/actor" // Todo service is not coupled to the auth service and just to the actor package
	"golang-practice/pkg/todoservice/internal/domain/entities/todo"
	"golang-practice/pkg/todoservice/internal/infrastructure/repositories"
	"golang-practice/pkg/todoservice/internal/permissions"

	"github.com/spf13/viper" // TODO We may want to avoid this coupling
)

type TodoService interface {
	AddTodo(createTodoDTO CreateTodoDTO, anActor actor.Actor) (TodoDTO, error)
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

func (state *todoServiceState) AddTodo(todoToCreate CreateTodoDTO, anActor actor.Actor) (TodoDTO, error) {

	var isAuthorized bool = permissions.Verify(anActor) // TODO find a way to enforce permissions automatically before every service's methods.

	if !isAuthorized {
		return emptyDTO, errors.New("actor is unauthorized")
	}

	var todoTask = todo.New(todo.TodoToCreate{
		Id:          todoToCreate.Id,
		Title:       todoToCreate.Title,
		Description: todoToCreate.Description,
		CreatedBy:   anActor.GetId(),
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
