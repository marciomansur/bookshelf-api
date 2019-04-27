package router

import (
	"net/http"

	handler "github.com/marciomansur/microservices-demo/bookshelf-api/pkg/handlers"
)

// Route contains routes metadata
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

// Routes is a list of routes, used in gorilla mux to create routing
type Routes []Route

var routes = Routes{
	Route{
		"GetBookshelf",
		"GET",
		"/bookshelf",
		handler.GetBookshelf,
	},
}
