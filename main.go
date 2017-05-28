package main

import (
	"log"
	"net/http"

	"github.com/Sandy987/go-api-exercise/routing"
	"github.com/Sandy987/go-api-exercise/todo"
)

func main() {

	router := routing.NewRouter(todo.Routes)

	log.Fatal(http.ListenAndServe(":8080", router))
}
