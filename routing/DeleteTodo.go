package routing

import (
	"errors"
	"net/http"
	"strconv"
	"todos/info"
	"todos/model"

	"github.com/ant0ine/go-json-rest/rest"
)

func DeleteTodo(w rest.ResponseWriter, r *rest.Request) {
	id, err := strconv.Atoi(r.PathParam("id"))
	if err != nil {
		rest.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = RemoveTodoByID(id)
	if err != nil {
		rest.Error(w, err.Error(), http.StatusBadRequest)
	}
	w.WriteHeader(http.StatusAccepted)
	w.WriteJson(map[string]string{"message": "Remove Todo Complete"})
}

func RemoveTodoByID(id int) error {
	todo, err := GetTodoByID(id)
	if err != nil {
		return err
	}
	index, err := FindArrayIndexByID(todo.ID)
	if err != nil {
		return err
	}
	info.Todos = remove(info.Todos, index)
	return nil
}

func FindArrayIndexByID(id int) (int, error) {
	for i := range info.Todos {
		if info.Todos[i].ID == id {
			return i, nil
		}
	}
	return 0, errors.New("ID not found")
}

func remove(slice []model.Todo, s int) []model.Todo {
	println("h")
	return append(slice[:s], slice[s+1:]...)
}
