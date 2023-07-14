package server

import (
	"golang-practice/pkg/env"
	"golang-practice/pkg/server/internal/controllers"

	"github.com/gin-gonic/gin"
)

func New(environment env.AppEnvironment) *gin.Engine {
	gin.SetMode(gin.ReleaseMode)

	var controller controllers.Controllers = controllers.New()
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
