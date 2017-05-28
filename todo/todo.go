package todo

import "time"

// Todo describes a task to be completed
type Todo struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Completed bool      `json:"completed"`
	Due       time.Time `json:"due"`
}

// Todos contains a collection of Todo items
type Todos []Todo
