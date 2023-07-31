package fakeAuth

import (
	"golang-practice/pkg/actor"
	"golang-practice/pkg/authservice"
	"golang-practice/pkg/server/internal/errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

// These middleware and methods make the link between the auth service and Gin without coupling them.
// Auth service doesn't know about Gin and can be used in another context.

func Middleware(c *gin.Context) {
	var currentActor, err = authservice.Authenticate(c.GetHeader("authorization"))

	if err != nil {
		c.JSON(http.StatusUnauthorized, errors.ServerError{Message: "Cannot authorize the request"})
		return
	}

	c.Set("actor", currentActor)
	c.Next()
}

func ExtractActorFromGinContext(c *gin.Context) actor.Actor {
	var actorValue, _ = c.Get("actor") // No need to check if it exists the middleware would have return an error otherwise.
	var currentActor actor.Actor = actorValue.(actor.Actor)
	return currentActor
}
