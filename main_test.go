package main

import (
	"bytes"
	"encoding/json"
	"golang-practice/pkg/children"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetChildByIdRoute(t *testing.T) {

	router := buildRouter("test")

	response := httptest.NewRecorder()
	request, _ := http.NewRequest("GET", "/children/1", nil)
	router.ServeHTTP(response, request)

	assert.Equal(t, 200, response.Code)

	result := children.Children{}
	if err := json.NewDecoder(response.Body).Decode(&result); err != nil {
		log.Fatalln(err)
	}

	assert.Equal(t, (result.Id), "1")
}

func TestGetChildrensRoute(t *testing.T) {
	router := buildRouter("test")

	response := httptest.NewRecorder()
	request, _ := http.NewRequest("GET", "/children", nil)
	router.ServeHTTP(response, request)

	assert.Equal(t, 200, response.Code)

	results := []children.Children{}
	if err := json.NewDecoder(response.Body).Decode(&results); err != nil {
		log.Fatalln(err)
	}

	assert.Equal(t, len(results), 3)
}

func TestAddChildRoute(t *testing.T) {
	router := buildRouter("test")

	childToCreate := children.Children{"5", "test", "test", "2023-12-23"}
	var childJSON, err = json.Marshal(childToCreate)

	if err != nil {
		return
	}

	response := httptest.NewRecorder()
	request, _ := http.NewRequest("POST", "/children", bytes.NewReader(childJSON))
	router.ServeHTTP(response, request)

	assert.Equal(t, http.StatusCreated, response.Code)

	result := children.Children{}

	if err := json.NewDecoder(response.Body).Decode(&result); err != nil {
		log.Fatalln(err)
	}

	assert.Equal(t, result.Id, childToCreate.Id)
}
