package main

import (
	"log"
	"net/http"
	"todo/route"

	"github.com/ant0ine/go-json-rest/rest"
)

func main() {
	api := rest.NewApi()
	api.Use(rest.DefaultDevStack...)
	router, err := rest.MakeRouter(
		rest.Get("/todos", route.GetTodos),
		rest.Get("/todo/:id", route.GetTodo),
		rest.Put("/todo/:id", route.PutTodo),
		rest.Post("/todo", route.PostTodo),
		rest.Delete("/todo/:id", route.DeleteTodo),
	)
	if err != nil {
		log.Fatal(err.Error())
	}
	api.SetApp(router)
	log.Fatal(http.ListenAndServe(":8080", api.MakeHandler()))
}
