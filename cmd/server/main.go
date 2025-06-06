package main

import (
	"log"
	"net/http"

	"github.com/medvedevse/quote-service/internal/router"
)

func main() {
	r := router.InitRouter()

	log.Fatal(http.ListenAndServe(":8080", r))
}
