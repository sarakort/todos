package main

import (
	"log"
	"net/http"
	"todos/routing"

	"github.com/ant0ine/go-json-rest/rest"
)

func main() {
	api := rest.NewApi()
	api.Use(rest.DefaultDevStack...)
	router, err := rest.MakeRouter(
		rest.Get("/todos", routing.GetTodos),
		rest.Get("/todo/:id", routing.GetTodoById),
		rest.Put("/todo/:id", routing.PutTodo),
		rest.Post("/todo", routing.PostTodo),
		rest.Delete("/todo/:id", routing.DeleteTodo),
	)
	if err != nil {
		log.Fatal(err.Error())
	}
	api.SetApp(router)
	log.Fatal(http.ListenAndServe(":8080", api.MakeHandler()))
}
