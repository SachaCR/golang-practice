package server

import (
	"golang-practice/pkg/env"
	"golang-practice/pkg/server/internals/controllers"
	"golang-practice/pkg/todo"

	"github.com/gin-gonic/gin"
)

type ServerDependencies struct {
	TodoRepository todo.TodoRepository
}

func New(environment env.AppEnvironment, dependencies ServerDependencies) *gin.Engine {
	gin.SetMode(gin.ReleaseMode)

	var ctrlDependencies = controllers.ControllerDependencies{
		TodoRepository: dependencies.TodoRepository,
	}

	var controller controllers.Controllers = controllers.New(ctrlDependencies)
	server := gin.New()

	if environment != env.Test {
		server.Use(gin.Logger())
		server.Use(gin.Recovery())
	}

	server.GET("/todos", controller.GetAllTodos)
	server.GET("/todos/:id", controller.GetTodoById)
	server.POST("/todos", controller.AddTodo)

	return server
}
