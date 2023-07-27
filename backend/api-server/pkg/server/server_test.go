package server

import (
	"bytes"
	"encoding/json"
	"golang-practice/pkg/env"
	"golang-practice/pkg/server/internal/errors"
	"golang-practice/pkg/todoservice"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

const TODO_PATH_PREFIX = "/todos"

func createATodo(t *testing.T, router *gin.Engine) {

	createTodoDTO := todoservice.CreateTodoDTO{
		Id:          "1",
		Title:       "test",
		Description: "test",
	}

	var childJSON, err = json.Marshal(createTodoDTO)

	if err != nil {
		t.Fail()
	}

	response := httptest.NewRecorder()
	request, _ := http.NewRequest("POST", TODO_PATH_PREFIX, bytes.NewReader(childJSON))
	router.ServeHTTP(response, request)

	assert.Equal(t, http.StatusCreated, response.Code)

	result := todoservice.TodoDTO{}
	if err := json.NewDecoder(response.Body).Decode(&result); err != nil {
		log.Fatalln(err)
	}

	assert.Equal(t, result.Id, createTodoDTO.Id)
}

func TestGetTodoByIdRoute(t *testing.T) {

	router := New(env.Test)

	createATodo(t, router)

	response := httptest.NewRecorder()
	request, _ := http.NewRequest("GET", TODO_PATH_PREFIX+"/1", nil)
	router.ServeHTTP(response, request)

	assert.Equal(t, 200, response.Code)

	result := todoservice.TodoDTO{}

	if err := json.NewDecoder(response.Body).Decode(&result); err != nil {
		log.Fatalln(err)
	}

	assert.Equal(t, (result.Id), "1")
}
func TestGetTodoByIdNotFoundRoute(t *testing.T) {

	router := New(env.Test)

	response := httptest.NewRecorder()
	request, _ := http.NewRequest("GET", TODO_PATH_PREFIX+"/4", nil)
	router.ServeHTTP(response, request)

	assert.Equal(t, 404, response.Code)

	result := errors.ServerError{}

	if err := json.NewDecoder(response.Body).Decode(&result); err != nil {
		log.Fatalln(err)
	}

	assert.Equal(t, result.Message, "Todo not found")
}

func TestGetTodosRoute(t *testing.T) {

	router := New(env.Test)

	createATodo(t, router)

	response := httptest.NewRecorder()
	request, _ := http.NewRequest("GET", TODO_PATH_PREFIX, nil)
	router.ServeHTTP(response, request)

	assert.Equal(t, 200, response.Code)

	results := []todoservice.TodoDTO{}
	if err := json.NewDecoder(response.Body).Decode(&results); err != nil {
		log.Fatalln(err)
	}

	assert.Equal(t, len(results), 1)
}

func TestAddTodoRoute(t *testing.T) {
	router := New(env.Test)

	createTodoDTO := todoservice.CreateTodoDTO{
		Id:          "5",
		Title:       "test",
		Description: "test",
	}

	var childJSON, err = json.Marshal(createTodoDTO)

	if err != nil {
		t.Fail()
	}

	response := httptest.NewRecorder()
	request, _ := http.NewRequest("POST", TODO_PATH_PREFIX, bytes.NewReader(childJSON))
	router.ServeHTTP(response, request)

	assert.Equal(t, http.StatusCreated, response.Code)

	result := todoservice.TodoDTO{}
	if err := json.NewDecoder(response.Body).Decode(&result); err != nil {
		log.Fatalln(err)
	}

	assert.Equal(t, result.Id, createTodoDTO.Id)
}

func TestAddTodoBadRequestRoute(t *testing.T) {
	router := New(env.Test)

	jsonStr := []byte("{ \"test\": \"Bad payload\"}")

	response := httptest.NewRecorder()
	request, _ := http.NewRequest("POST", TODO_PATH_PREFIX, bytes.NewBuffer(jsonStr))

	router.ServeHTTP(response, request)

	assert.Equal(t, http.StatusBadRequest, response.Code)

	result := errors.ServerError{}

	if err := json.NewDecoder(response.Body).Decode(&result); err != nil {
		log.Fatalln(err)
	}

	assert.Equal(t, result.Message, "Cannot bind payload")
}
