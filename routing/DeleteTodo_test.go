package routing

import "testing"

func TestRemoveTodo(t *testing.T) {
	id := 2
	removeTodo(2)

	todo, err := GetTodoByID(2)
	if err != nil {
		t.Error(err.Error())
	}

}
