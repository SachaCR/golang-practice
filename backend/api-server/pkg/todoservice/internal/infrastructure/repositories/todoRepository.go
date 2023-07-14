package repositories

import (
	"fmt"
	"golang-practice/pkg/todoservice/internal/domain/entities/todo"
	"golang-practice/pkg/todoservice/internal/infrastructure/repositories/internal/inmemory"
	"golang-practice/pkg/todoservice/internal/infrastructure/repositories/internal/postgres"
)

func New(engine RepositoryEngine) todo.TodoRepository {

	switch engine.String() {
	case "Postgres":
		return postgres.New()

	case "InMemory":
		return inmemory.New()
	}

	fmt.Println("WARNING: todo repository unknown repository engine. Default InMemory engine used")
	return inmemory.New()
}
