package route

import "github.com/ant0ine/go-json-rest/rest"

func GetTodos(w rest.ResponseWriter, r *rest.Request) {
	w.WriteJson(nil)
}
