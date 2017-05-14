package routing

import (
	"net/http"
	"todos/info"
	"todos/model"

	"github.com/ant0ine/go-json-rest/rest"
)

func PostTodo(w rest.ResponseWriter, r *rest.Request) {
	todo := model.Todo{}
	err := r.DecodeJsonPayload(&todo)
	if err != nil {
		rest.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	id := CreateNewTodo(todo)
	w.WriteHeader(http.StatusCreated)
	w.WriteJson(map[string]int{"ID": id})
}

func CreateNewTodo(todo model.Todo) int {
	info.Count++
	todo.ID = info.Count
	info.Todos = append(info.Todos, todo)
	return GetTodoByID(todo.ID).ID
}
