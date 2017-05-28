package main

import (
	"Sandy987/go-api-exercise/routing"
	"Sandy987/go-api-exercise/todo"
	"log"
	"net/http"
)

func main() {

	router := routing.NewRouter(todo.Routes)

	log.Fatal(http.ListenAndServe(":8080", router))
}
