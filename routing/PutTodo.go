package routing

import (
	"net/http"
	"todos/info"
	"todos/model"

	"strconv"

	"github.com/ant0ine/go-json-rest/rest"
)

func PutTodo(w rest.ResponseWriter, r *rest.Request) {
	id, err := strconv.Atoi(r.PathParam("id"))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	todo := GetTodoByID(id)
	err = r.DecodeJsonPayload(&todo)
	result := UpdateTodo(id, todo)
	w.WriteJson(map[string]int{"ID": result.ID})
	w.WriteHeader(http.StatusAccepted)

}

func UpdateTodo(id int, todo model.Todo) model.Todo {
	for i := range info.Todos {
		if info.Todos[i].ID == id {
			info.Todos[i] = todo
			info.Todos[i].ID = id
			return info.Todos[i]
		}
	}

	return model.Todo{}
}
