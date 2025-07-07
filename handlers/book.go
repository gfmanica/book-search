package handlers

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/gfmanica/book-search/services"
)

type EnrichRequest struct {
	ISBN string `json:"isbn"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}

func EnrichBookHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)

		json.NewEncoder(w).Encode(ErrorResponse{Error: "Método não permitido"})

		return
	}

	body, err := io.ReadAll(r.Body)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)

		json.NewEncoder(w).Encode(ErrorResponse{Error: "Erro ao ler o corpo da requisição"})

		return
	}

	defer r.Body.Close()

	var req EnrichRequest

	if err := json.Unmarshal(body, &req); err != nil || req.ISBN == "" {
		w.WriteHeader(http.StatusBadRequest)

		json.NewEncoder(w).Encode(ErrorResponse{Error: "JSON inválido ou ISBN ausente"})

		return
	}

	book, err := services.FetchBookFromGoogle(req.ISBN)

	if err != nil {
		w.WriteHeader(http.StatusNotFound)

		json.NewEncoder(w).Encode(ErrorResponse{Error: err.Error()})

		return
	}

	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(book)
}
