package todo

import "Sandy987/go-api-exercise/routing"

// Routes contains all potential handlers for the Todo module
var Routes = routing.Routes{
	routing.Route{
		Name:        "Index",
		Method:      "GET",
		Pattern:     "/",
		HandlerFunc: Index,
	},
	routing.Route{
		Name:        "List",
		Method:      "GET",
		Pattern:     "/todos",
		HandlerFunc: List,
	},
	routing.Route{
		Name:        "Show",
		Method:      "GET",
		Pattern:     "/todos/{todoId}",
		HandlerFunc: Show,
	},
	routing.Route{
		Name:        "Make",
		Method:      "POST",
		Pattern:     "/todos",
		HandlerFunc: Make,
	},
	routing.Route{
		Name:        "Delete",
		Method:      "DELETE",
		Pattern:     "todos/{todoId}",
		HandlerFunc: Delete,
	},
}
