package main

import (
	"log"
	"net/http"
	"training/todos/handler"
)

func main() {
	api := handler.NewApi()
	log.Fatal(http.ListenAndServe(":8080", api.MakeHandler()))
}
