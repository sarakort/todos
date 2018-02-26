package routing

import (
	"errors"
	"net/http"
	"training/todos/info"
	"training/todos/model"

	"strconv"

	"github.com/ant0ine/go-json-rest/rest"
)

func GetTodo(w rest.ResponseWriter, r *rest.Request) {
	id, err := strconv.Atoi(r.PathParam("id"))
	if err != nil {
		rest.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	todo, err := GetTodoByID(id)
	if err != nil {
		rest.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteJson(todo)
}

func GetTodoByID(id int) (model.Todo, error) {
	todo := model.Todo{}
	for _, v := range info.Todos {
		if v.ID == id {
			todo = v
			return todo, nil
		}
	}
	return todo, errors.New("Id not found")

}
