package todo

import "fmt"

var currentID int

var todoStore Todos

// GetAllTodos returns all available todo items
func GetAllTodos() Todos {
	return todoStore
}

// GetTodo retrieves a Todo with the given ID
func GetTodo(id int) Todo {
	for _, t := range todoStore {
		if t.ID == id {
			return t
		}
	}
	// return empty Todo if not found
	return Todo{}
}

// CreateTodo create a Todo object in the repository
func CreateTodo(t Todo) Todo {
	currentID++
	t.ID = currentID
	todoStore = append(todoStore, t)
	return t
}

// DeleteTodo removes a Todo with the given ID from the repository
func DeleteTodo(id int) error {
	for i, t := range todoStore {
		if t.ID == id {
			todoStore = append(todoStore[:i], todoStore[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("Could not find Todo with id of %d to delete", id)
}
