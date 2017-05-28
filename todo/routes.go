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
		Name:        "TodoIndex",
		Method:      "GET",
		Pattern:     "/todos",
		HandlerFunc: List,
	},
	routing.Route{
		Name:        "TodoShow",
		Method:      "GET",
		Pattern:     "/todos/{todoId}",
		HandlerFunc: Show,
	},
}
