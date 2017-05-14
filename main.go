package main

import (
	"log"

	"net/http"

	"github.com/ant0ine/go-json-rest/rest"
)

func main() {
	api := rest.NewApi()
	api.Use(rest.DefaultDevStack...)
	router, err := rest.MakeRouter(
		rest.Get("/todos", GetTodos),
		rest.Get("/todo/:id", GetTodo),
		rest.Put("/todo/:id", PutTodo),
		rest.Post("/todo", PostTodo),
		rest.Delete("/todo/:id", DeleteTodo),
	)
	if err != nil {
		log.Fatal(err.Error())
	}
	api.SetApp(router)
	log.Fatal(http.ListenAndServe(":8080", api.MakeHandler()))
}

func GetTodos(w rest.ResponseWriter, r *rest.Request) {
	w.WriteJson(nil)
}

func GetTodo(w rest.ResponseWriter, r *rest.Request) {

}

func PostTodo(w rest.ResponseWriter, r *rest.Request) {

}

func DeleteTodo(w rest.ResponseWriter, r *rest.Request) {

}

func PutTodo(w rest.ResponseWriter, r *rest.Request) {

}
