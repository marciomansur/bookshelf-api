package handlers

import (
	"encoding/json"
	"net/http"

	books "github.com/marciomansur/microservices-demo/bookshelf-api/internal/models"
)

type IDParam struct {
	// required: true
	ID int64 `json:"id"`
}

func GetBookshelf(w http.ResponseWriter, r *http.Request) {
	// header
	w.Header().Set("Content type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(books.Get())
}
