package routing

import (
	"testing"
	"todos/model"
)

func TestCreateNewTodo(t *testing.T) {
	var todo = model.Todo{Message: "Homework 3"}

	id := CreateNewTodo(todo)
	todo = GetTodoByID(id)

	if todo.ID == 0 {
		t.Error("Not Found Id")
	}

}
