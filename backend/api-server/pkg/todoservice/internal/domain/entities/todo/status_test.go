package todo

import (
	"testing"

	"github.com/go-playground/assert/v2"
)

func TestStatusToString(t *testing.T) {
	var todoStatus = Todo
	var doneStatus = Done
	var inProgressStatus = InProgress
	var cancelledStatus = Cancelled
	var unknownStatus = Unknown

	assert.Equal(t, todoStatus.String(), "Todo")
	assert.Equal(t, doneStatus.String(), "Done")
	assert.Equal(t, inProgressStatus.String(), "InProgress")
	assert.Equal(t, cancelledStatus.String(), "Cancelled")
	assert.Equal(t, unknownStatus.String(), "Unknown")
}

func TestStatusFromString(t *testing.T) {
	var todoStatus = TodoStatusFromString("Todo")
	var doneStatus = TodoStatusFromString("Done")
	var inProgressStatus = TodoStatusFromString("InProgress")
	var cancelledStatus = TodoStatusFromString("Cancelled")
	var unknownStatus = TodoStatusFromString("Unknown")

	assert.Equal(t, todoStatus, Todo)
	assert.Equal(t, doneStatus, Done)
	assert.Equal(t, inProgressStatus, InProgress)
	assert.Equal(t, cancelledStatus, Cancelled)
	assert.Equal(t, unknownStatus, Unknown)
}
