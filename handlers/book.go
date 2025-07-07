package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/gfmanica/book-search/services"
)

type ErrorResponse struct {
	Error string `json:"error"`
}

func EnrichBookHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)

		json.NewEncoder(w).Encode(ErrorResponse{Error: "Método não permitido. Use GET."})

		return
	}

	isbn := r.URL.Query().Get("isbn")

	if isbn == "" {
		w.WriteHeader(http.StatusBadRequest)

		json.NewEncoder(w).Encode(ErrorResponse{Error: "Parâmetro 'isbn' é obrigatório na query string."})

		return
	}

	book, err := services.FetchBookFromGoogle(isbn)

	if err != nil {
		w.WriteHeader(http.StatusNotFound)

		json.NewEncoder(w).Encode(ErrorResponse{Error: err.Error()})

		return
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(book)
}
