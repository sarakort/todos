package routing

import (
	"strconv"
	"testing"
	"todos/model"
)

func TestUpdateTodo(t *testing.T) {
	todo := model.Todo{Message: "Update", Checked: true}
	id := 2
	result := UpdateTodo(id, todo)

	if id != result.ID {
		t.Error("Not Equal, Expect:" + strconv.Itoa(id) + " but got " + strconv.Itoa(result.ID))
	}

}
