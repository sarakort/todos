package routing

import (
	"testing"
	"todos/model"
)

func TestCreateNewTodo(t *testing.T) {
	var todo = model.Todo{Message: "Homework 3"}

	id := CreateNewTodo(todo)
	todo, err := GetTodoByID(id)
	if err != nil {
		t.Error(err.Error())
	}

	if todo.ID == 0 {
		t.Error("Not Found Id")
	}

}
