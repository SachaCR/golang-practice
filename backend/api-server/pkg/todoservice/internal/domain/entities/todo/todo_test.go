package todo

import (
	"testing"
	"time"

	"github.com/go-playground/assert/v2"
)

func TestNew(t *testing.T) {
	var dateFormatString = "2006-01-02"
	var todayDateString = time.Now()
	var todoToCreate = TodoToCreate{Id: "1", Title: "Task Title", Description: "Task Description"}
	var myTodo = New(todoToCreate)

	var todoTaskState = myTodo.CopyState()

	assert.Equal(t, todoTaskState.Id, "1")
	assert.Equal(t, todoTaskState.Description, "Task Description")
	assert.Equal(t, todoTaskState.Title, "Task Title")
	assert.Equal(t, todoTaskState.Status, Todo)
	assert.Equal(t, todoTaskState.CreatedAt.Format(dateFormatString), todayDateString.Format(dateFormatString))
}
