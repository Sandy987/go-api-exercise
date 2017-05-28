package todo

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"io"
	"io/ioutil"

	"github.com/gorilla/mux"
)

// Index gives an introduction to the Todo module
func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome to the Todo module!")
}

// List gives a list of all available Todos
func List(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(GetAllTodos()); err != nil {
		panic(err)
	}
}

// Show shows the details of a particular Todo
func Show(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	todoID, converr := strconv.Atoi(vars["todoId"])
	if converr != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(GetTodo(todoID)); err != nil {
		panic(err)
	}
}

// Make accepts a JSON object and creates a Todo with it
func Make(w http.ResponseWriter, r *http.Request) {
	var todo Todo
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}
	if err := r.Body.Close(); err != nil {
		panic(err)
	}

	if err := json.Unmarshal(body, &todo); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusBadRequest)
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}

	result := CreateTodo(todo)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(result); err != nil {
		panic(err)
	}
}

// Delete removes a TODO using an ID
func Delete(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	todoID, converr := strconv.Atoi(vars["todoId"])
	if converr != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err := DeleteTodo(todoID)
	if err != nil {
		panic(err)
	}

	w.WriteHeader(http.StatusOK)
}
