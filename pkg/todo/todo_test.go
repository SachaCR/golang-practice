package todo

import (
	"testing"
	"time"

	"github.com/go-playground/assert/v2"
)

func TestNew(t *testing.T) {
	var todayDateString = time.Now().Format("2006-01-02")
	var todoToCreate = TodoToCreate{Id: "1", Title: "Task Title", Description: "Task Description"}
	var myTodo = New(todoToCreate)

	assert.Equal(t, myTodo.ToDTO().Id, "1")
	assert.Equal(t, myTodo.ToDTO().Description, "Task Description")
	assert.Equal(t, myTodo.ToDTO().Title, "Task Title")
	assert.Equal(t, myTodo.ToDTO().Status, "Todo")
	assert.Equal(t, myTodo.ToDTO().CreatedAt, todayDateString)
}

func TestFromDTO(t *testing.T) {
	var todoDTO = TodoDTO{Id: "1", Title: "Task Title", Description: "Task Description", CreatedAt: "2022-12-23", Status: "InProgress"}
	var myTodo, err = FromDTO(todoDTO)

	if err != nil {
		t.Fail()
	}

	assert.Equal(t, myTodo.ToDTO().Id, "1")
	assert.Equal(t, myTodo.ToDTO().Description, "Task Description")
	assert.Equal(t, myTodo.ToDTO().Title, "Task Title")
	assert.Equal(t, myTodo.ToDTO().Status, "InProgress")
	assert.Equal(t, myTodo.ToDTO().CreatedAt, "2022-12-23")
}
