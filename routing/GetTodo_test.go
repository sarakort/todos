package routing

import (
	"strconv"
	"testing"
	"todos/model"
)

func TestGetTodoAndExpectingId(t *testing.T) {
	id := 2

	todo := GetTodoByID(id)
	validateResult(t, id, todo)

}

func validateResult(t *testing.T, expecting_id int, result model.Todo) {
	if result.ID != expecting_id {
		t.Error("expect '" + strconv.Itoa(expecting_id) + "' but got " + strconv.Itoa(result.ID))
	}
}
