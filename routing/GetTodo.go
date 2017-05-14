package routing

import (
	"net/http"
	"todos/info"
	"todos/model"

	"strconv"

	"github.com/ant0ine/go-json-rest/rest"
)

func GetTodo(w rest.ResponseWriter, r *rest.Request) {
	id, err := strconv.Atoi(r.PathParam("id"))
	if err != nil {
		rest.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteJson(GetTodoByID(id))
}

func GetTodoByID(id int) model.Todo {
	todo := model.Todo{}
	for _, v := range info.Todos {
		if v.ID == id {
			todo = v
			break
		}
	}
	return todo

}
