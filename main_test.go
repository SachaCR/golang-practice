package main

import (
	"bytes"
	"encoding/json"
	"golang-practice/pkg/baby"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetBabyByIdRoute(t *testing.T) {
	router := buildRouter()

	response := httptest.NewRecorder()
	request, _ := http.NewRequest("GET", "/babies/1", nil)
	router.ServeHTTP(response, request)

	assert.Equal(t, 200, response.Code)

	result := baby.Baby{}
	if err := json.NewDecoder(response.Body).Decode(&result); err != nil {
		log.Fatalln(err)
	}

	assert.Equal(t, (result.Id), "1")
}

func TestGetBabiesRoute(t *testing.T) {
	router := buildRouter()

	response := httptest.NewRecorder()
	request, _ := http.NewRequest("GET", "/babies", nil)
	router.ServeHTTP(response, request)

	assert.Equal(t, 200, response.Code)

	results := []baby.Baby{}
	if err := json.NewDecoder(response.Body).Decode(&results); err != nil {
		log.Fatalln(err)
	}

	assert.Equal(t, len(results), 3)
}

func TestAddBabyRoute(t *testing.T) {
	router := buildRouter()

	babyToCreate := &baby.Baby{Id: "5", FirstName: "test", LastName: "test", BirthDate: "2023-12-23"}
	var babyJSON, err = json.Marshal(babyToCreate)

	if err != nil {
		return
	}

	response := httptest.NewRecorder()
	request, _ := http.NewRequest("POST", "/babies", bytes.NewReader(babyJSON))
	router.ServeHTTP(response, request)

	assert.Equal(t, http.StatusCreated, response.Code)

	result := baby.Baby{}
	if err := json.NewDecoder(response.Body).Decode(&result); err != nil {
		log.Fatalln(err)
	}

	assert.Equal(t, result.Id, babyToCreate.Id)
}
