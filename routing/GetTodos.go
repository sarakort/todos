package routing

import (
	"todo/info"

	"github.com/ant0ine/go-json-rest/rest"
)

func GetTodos(w rest.ResponseWriter, r *rest.Request) {
	w.WriteJson(info.Todos)
}
