package authservice

import "golang-practice/pkg/actor"

// This service is not coupled to Gin it just knows the app need an actor to work with.
type AuthenticationService interface {
	Authenticate(authHeader string) actor.Actor
}

func Authenticate(authHeader string) (actor.Actor, error) {
	// Verify JWT, verify signature, or whatever to build an Actor object
	var currentActor actor.Actor = actor.New("toto", actor.User, []string{"basic-user"})
	return currentActor, nil
}
