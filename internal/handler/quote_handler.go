package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"

	"quote-service/internal/model"
	"quote-service/internal/repository"

	"github.com/gorilla/mux"
)

var quoteRepo = repository.QuoteRepository()

func GetQuotes(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if len(r.URL.RawQuery) > 0 {
		authorName := r.URL.Query().Get("author")
		if authorName == "" {
			http.Error(w, "Author's name is empty", http.StatusBadRequest)
			return
		}

		var sortedQuotes []model.Quote

		for _, quote := range quoteRepo {
			if quote.Author == authorName {
				sortedQuotes = append(sortedQuotes, quote)
			}
		}

		err := json.NewEncoder(w).Encode(sortedQuotes)
		if err != nil {
			log.Println("Ошибка получения списка цитат от автора")
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		return
	}

	err := json.NewEncoder(w).Encode(quoteRepo)
	if err != nil {
		log.Println("Ошибка получения списка цитат")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func AddQuote(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var quote model.Quote
	var err error

	err = json.NewDecoder(r.Body).Decode(&quote)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	quote.ID = fmt.Sprintf("%d", len(quoteRepo)+1)
	quoteRepo = append(quoteRepo, quote)
	err = json.NewEncoder(w).Encode(quote)
	if err != nil {
		log.Println("Ошибка при добавлении цитаты")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func DeleteQuote(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for i, quote := range quoteRepo {
		if quote.ID == params["id"] {
			quoteRepo = append(quoteRepo[:i], quoteRepo[i+1:]...)
			break
		}
	}

	err := json.NewEncoder(w).Encode(quoteRepo)
	if err != nil {
		log.Println("Ошибка при удалении цитаты")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func GetRandomQuote(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	randomQuote := quoteRepo[rand.Intn(len(quoteRepo))]
	err := json.NewEncoder(w).Encode(randomQuote)
	if err != nil {
		log.Println("Ошибка получения рандомной цитаты")
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	message := "1. Добавление новой цитаты (POST /quotes)\n\n2. Получение всех цитат (GET /quotes)\n\n3. Получение случайной цитаты (GET /quotes/random)\n\n4. Фильтрация по автору (GET /quotes?author=Confucius)\n\n5. Удаление цитаты по ID (DELETE /quotes/{id})"
	fmt.Fprint(w, message)
}
