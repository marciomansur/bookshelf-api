package main

import (
	"log"
	"net/http"

	model "github.com/marciomansur/bookshelf-api/models"
	"github.com/marciomansur/bookshelf-api/router"
	"github.com/rs/cors"
)

func setGlobalMiddleware(handler http.Handler) http.Handler {
	handleCORS := cors.Default().Handler
	return handleCORS(handler)
}

func main() {
	model.Add(model.Book{ID: "1", Title: "The Kubernetes Book", Description: "A book about Kubernetes", Author: &model.Author{Name: "Jhonas"}})
	model.Add(model.Book{ID: "2", Title: "The Docker Book", Description: "A book about Docker", Author: &model.Author{Name: "Thomas"}})
	model.Add(model.Book{ID: "3", Title: "The Go Book", Description: "A book about Go", Author: &model.Author{Name: "George"}})

	router := router.NewRouter()
	log.Print(http.ListenAndServe(":8080", setGlobalMiddleware(router)))
}