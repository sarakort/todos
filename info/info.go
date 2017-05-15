package info

import "todos/model"

var todo1 = model.Todo{
	ID:      1,
	Checked: true,
	Message: "Homework 1",
}
var todo2 = model.Todo{
	ID:      2,
	Checked: false,
	Message: "Homework 2",
}
var Count int = 2
var Todos = []model.Todo{todo1, todo2}
