package controllers

import (
	"golang-practice/pkg/server/internals/errors"
	"golang-practice/pkg/todo"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ControllerDependencies struct {
	TodoRepository todo.TodoRepository
}

type Controllers interface {
	GetAllTodos(c *gin.Context)
	AddTodo(c *gin.Context)
	GetTodoById(c *gin.Context)
}

type ControllerState struct {
	todoRepository todo.TodoRepository
}

func (state *ControllerState) GetAllTodos(c *gin.Context) {
	var todos, err = state.todoRepository.GetAll()

	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			errors.ServerError{Message: err.Error()})
	}

	c.IndentedJSON(http.StatusOK, todos)
}

func (state *ControllerState) AddTodo(c *gin.Context) {
	var todoToCreate todo.TodoToCreate

	bindJsonError := c.BindJSON(&todoToCreate)

	if bindJsonError != nil {
		c.JSON(http.StatusBadRequest, errors.ServerError{Message: "Cannot bind payload"})
	}

	var todoCreated = todo.New(todoToCreate)

	state.todoRepository.Save(todoCreated)

	c.IndentedJSON(http.StatusCreated, todoCreated)
}

func (state *ControllerState) GetTodoById(c *gin.Context) {
	id := c.Param("id")

	var todo, err = state.todoRepository.GetById(id)

	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			errors.ServerError{Message: err.Error()})
	}

	if todo == nil {
		c.IndentedJSON(http.StatusNotFound, errors.ServerError{Message: "Todo not found"})
	}

	c.IndentedJSON(http.StatusOK, todo)
}

func New(dependencies ControllerDependencies) Controllers {
	return &ControllerState{
		todoRepository: dependencies.TodoRepository,
	}
}
