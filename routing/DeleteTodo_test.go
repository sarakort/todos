package routing

import "testing"

func TestRemoveTodo(t *testing.T) {
	id := 2
	err := RemoveTodoByID(id)
	if err != nil {
		t.Error(err.Error())
	}

}
