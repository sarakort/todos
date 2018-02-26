package routing

import (
	"training/todos/info"
	"training/todos/model"

	"github.com/ant0ine/go-json-rest/rest"
)

func GetTodos(w rest.ResponseWriter, r *rest.Request) {
	w.WriteJson(GetAllTodos())
}
func GetAllTodos() []model.Todo {
	return info.Todos
}
