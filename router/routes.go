package router

import (
	"net/http"

	handler "github.com/marciomansur/bookshelf-api/handlers"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		"GetBookshelf",
		"GET",
		"/bookshelf",
		handler.GetBookshelf,
	},
}
