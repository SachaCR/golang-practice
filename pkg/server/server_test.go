package server

import (
	"bytes"
	"encoding/json"
	"golang-practice/pkg/env"
	inmemory "golang-practice/pkg/repositories/inMemory"
	"golang-practice/pkg/server/internals/errors"
	"golang-practice/pkg/todo"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetTodoByIdRoute(t *testing.T) {

	todoRepository := inmemory.New()

	todoRepository.Save(todo.New(todo.TodoToCreate{
		Id:          "1",
		Title:       "My Todo Title",
		Description: "My Todo Description",
	}))

	router := New(env.Test, ServerDependencies{todoRepository})

	response := httptest.NewRecorder()
	request, _ := http.NewRequest("GET", "/todos/1", nil)
	router.ServeHTTP(response, request)

	assert.Equal(t, 200, response.Code)

	result := todo.TodoDTO{}

	if err := json.NewDecoder(response.Body).Decode(&result); err != nil {
		log.Fatalln(err)
	}

	assert.Equal(t, (result.Id), "1")
}
func TestGetTodoByIdNotFoundRoute(t *testing.T) {

	todoRepository := inmemory.New()
	router := New(env.Test, ServerDependencies{todoRepository})

	response := httptest.NewRecorder()
	request, _ := http.NewRequest("GET", "/todos/4", nil)
	router.ServeHTTP(response, request)

	assert.Equal(t, 404, response.Code)

	result := errors.ServerError{}
	if err := json.NewDecoder(response.Body).Decode(&result); err != nil {
		log.Fatalln(err)
	}

	assert.Equal(t, result.Message, "Todo not found")
}

func TestGetTodosRoute(t *testing.T) {
	todoRepository := inmemory.New()

	todoRepository.Save(todo.New(todo.TodoToCreate{
		Id:          "1",
		Title:       "My Todo Title",
		Description: "My Todo Description",
	}))

	router := New(env.Test, ServerDependencies{todoRepository})

	response := httptest.NewRecorder()
	request, _ := http.NewRequest("GET", "/todos", nil)
	router.ServeHTTP(response, request)

	assert.Equal(t, 200, response.Code)

	results := []todo.TodoDTO{}
	if err := json.NewDecoder(response.Body).Decode(&results); err != nil {
		log.Fatalln(err)
	}

	assert.Equal(t, len(results), 1)
}

func TestAddTodoRoute(t *testing.T) {
	todoRepository := inmemory.New()
	router := New(env.Test, ServerDependencies{todoRepository})

	todoToCreate := todo.TodoDTO{
		Id:          "5",
		Title:       "test",
		Description: "test",
		CreatedAt:   "2023-12-23",
	}

	var childJSON, err = json.Marshal(todoToCreate)

	if err != nil {
		t.Fail()
	}

	response := httptest.NewRecorder()
	request, _ := http.NewRequest("POST", "/todos", bytes.NewReader(childJSON))
	router.ServeHTTP(response, request)

	assert.Equal(t, http.StatusCreated, response.Code)

	result := todo.TodoDTO{}
	if err := json.NewDecoder(response.Body).Decode(&result); err != nil {
		log.Fatalln(err)
	}

	assert.Equal(t, result.Id, todoToCreate.Id)
}

func TestAddTodoBadRequestRoute(t *testing.T) {
	todoRepository := inmemory.New()
	router := New(env.Test, ServerDependencies{todoRepository})

	jsonStr := []byte("{ \"test\": \"Bad payload\"}")

	response := httptest.NewRecorder()
	request, _ := http.NewRequest("POST", "/todos", bytes.NewBuffer(jsonStr))

	router.ServeHTTP(response, request)

	assert.Equal(t, http.StatusBadRequest, response.Code)

	result := errors.ServerError{}

	if err := json.NewDecoder(response.Body).Decode(&result); err != nil {
		log.Fatalln(err)
	}

	assert.Equal(t, result.Message, "Cannot bind payload")
}
