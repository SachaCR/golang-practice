package todoservice

import (
	"testing"
	"time"

	"github.com/go-playground/assert/v2"
)

func TestFromDTO(t *testing.T) {
	var todoDTO = TodoDTO{Id: "1", Title: "Task Title", Description: "Task Description", CreatedAt: "2022-12-23", Status: "InProgress"}
	var myTodo, err = FromDTO(todoDTO)

	if err != nil {
		t.Fail()
	}

	assert.Equal(t, ToDTO(myTodo).Id, "1")
	assert.Equal(t, ToDTO(myTodo).Description, "Task Description")
	assert.Equal(t, ToDTO(myTodo).Title, "Task Title")
	assert.Equal(t, ToDTO(myTodo).Status, "InProgress")
	assert.Equal(t, ToDTO(myTodo).CreatedAt, "2022-12-23")
}

func TestFromDTOError(t *testing.T) {
	var todoDto TodoDTO = TodoDTO{Id: "1",
		Title:       "Title",
		Description: "Description",
		CreatedAt:   "dsajfaiosdj",
		Status:      "Todo",
	}

	var _, err = FromDTO(todoDto)

	var _, isParseError = err.(*time.ParseError)
	var errorMessage = err.Error()

	assert.Equal(t, isParseError, true)
	assert.Equal(t, errorMessage, "parsing time \"dsajfaiosdj\" as \"2006-01-02\": cannot parse \"dsajfaiosdj\" as \"2006\"")
}
