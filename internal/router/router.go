package router

import (
	"quote-service/internal/handler"

	"github.com/gorilla/mux"
)

func InitRouter() *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/", handler.HomeHandler).Methods("GET")
	r.HandleFunc("/quotes", handler.AddQuote).Methods("POST")
	r.HandleFunc("/quotes", handler.GetQuotes).Methods("GET")
	r.HandleFunc("/quotes/random", handler.GetRandomQuote).Methods("GET")
	r.HandleFunc("/quotes/{id}", handler.DeleteQuote).Methods("DELETE")

	return r
}
