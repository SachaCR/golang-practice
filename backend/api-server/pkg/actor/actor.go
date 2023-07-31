package actor

import "fmt"

// This package provide interfaces and types to represent the different actors that interact with our app.
// It is not tied to any auth or permissions mechanism.

type ActorType int64

const (
	User ActorType = iota
	Service
	Cron
	Job
)

type Actor interface {
	GetId() string
	GetType() ActorType
	GetRoles() []string
}

type actorState struct {
	Id        string
	ActorType ActorType
	Roles     []string
}

func New(id string, actorType ActorType, roles []string) Actor {
	return &actorState{
		Id:        id,
		ActorType: actorType,
		Roles:     roles,
	}
}

func (state *actorState) GetId() string {
	return state.Id
}

func (state *actorState) GetType() ActorType {
	return state.ActorType
}

func (state *actorState) GetRoles() []string {
	return state.Roles
}

func (s ActorType) String() string {
	switch s {

	case User:
		return "user"

	case Service:
		return "service"

	case Cron:
		return "cron"

	case Job:
		return "job"
	}

	return "unknown"
}

func ParseActorTypeFromString(actorString string) ActorType {
	switch actorString {

	case "user":
		return User

	case "service":
		return Service

	case "cron":
		return Cron

	case "job":
		return Job
	}

	panic(fmt.Errorf("UNKNOWN ACTOR TYPE: " + actorString))
}
