package handler

import (
	"log"
	"training/todos/routing"

	"github.com/ant0ine/go-json-rest/rest"
)

func NewApi() *rest.Api {
	api := rest.NewApi()
	api.Use(rest.DefaultDevStack...)
	api.Use(&rest.CorsMiddleware{
		RejectNonCorsRequests: false,
		AllowedMethods: []string{"GET", "POST", "PUT" ,"DELETE" },
		AllowedHeaders: []string{
			"Accept", "Content-Type", "X-Custom-Header", "Origin"},
		AccessControlAllowCredentials: true,
		AccessControlMaxAge:           3600,
	})
	router, err := rest.MakeRouter(
		rest.Get("/todos", routing.GetTodos),
		rest.Get("/todos/:id", routing.GetTodo),
		rest.Put("/todos/:id", routing.PutTodo),
		rest.Post("/todos", routing.PostTodo),
		rest.Delete("/todos/:id", routing.DeleteTodo),
	)
	if err != nil {
		log.Fatal(err.Error())
	}
	api.SetApp(router)
	return api
}
