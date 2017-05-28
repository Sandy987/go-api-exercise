package todo

import "testing"
import "time"

func TestCreateTodo(t *testing.T) {
	todo := Todo{
		Name:      "MyTodo",
		Completed: false,
		Due:       time.Now(),
	}
	result := CreateTodo(todo)
	if result.ID == 0 {
		t.Error("Expected the TODO to have an ID, got ", result.ID)
	}
}

func TestGetTodo(t *testing.T) {
	todo := Todo{
		Name:      "MyTodo",
		Completed: false,
		Due:       time.Now(),
	}
	created := CreateTodo(todo)
	retrieved := GetTodo(created.ID)
	if retrieved.ID != created.ID {
		t.Error("Expected to retrieve created TODO")
	}
}
