package main

import (
	"log"
	"net/http"

	"github.com/gfmanica/book-search/handlers"
)

func main() {
	http.HandleFunc("/enrich", handlers.EnrichBookHandler)

	log.Println("Servidor iniciado na porta 8080...")

	log.Fatal(http.ListenAndServe(":8080", nil))
}
