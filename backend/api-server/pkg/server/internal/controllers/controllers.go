package controllers

import (
	"golang-practice/pkg/server/internal/errors"
	"golang-practice/pkg/todoservice"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Controllers interface {
	GetAllTodos(c *gin.Context)
	AddTodo(c *gin.Context)
	GetTodoById(c *gin.Context)
}

type ControllerState struct {
	TodoService todoservice.TodoService
}

func New() Controllers {
	return &ControllerState{
		TodoService: todoservice.New(),
	}
}

func (state *ControllerState) GetAllTodos(c *gin.Context) {
	var todos, err = state.TodoService.GetTodoList()

	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			errors.ServerError{Message: err.Error()})
	}

	c.IndentedJSON(http.StatusOK, todos)
}

func (state *ControllerState) AddTodo(c *gin.Context) {
	var createTodoDTO todoservice.CreateTodoDTO

	bindJsonError := c.BindJSON(&createTodoDTO)

	if bindJsonError != nil {
		c.JSON(http.StatusBadRequest, errors.ServerError{Message: "Cannot bind payload"})
		return
	}

	var todoDTO, err = state.TodoService.AddTodo(createTodoDTO)

	if err != nil {
		c.JSON(http.StatusInternalServerError, errors.ServerError{Message: "Cannot create todo"})
	}

	c.IndentedJSON(http.StatusCreated, todoDTO)
}

func (state *ControllerState) GetTodoById(c *gin.Context) {
	id := c.Param("id")

	var todoDTO, err = state.TodoService.GetTodoById(id)

	if err != nil {

		var _, isNotFoundError = err.(*todoservice.TodoNotFoundError)

		if isNotFoundError {
			c.IndentedJSON(http.StatusNotFound, errors.ServerError{Message: err.Error()})
			return
		}

		c.JSON(
			http.StatusInternalServerError,
			errors.ServerError{Message: err.Error()})
		return
	}

	c.IndentedJSON(http.StatusOK, todoDTO)
}
