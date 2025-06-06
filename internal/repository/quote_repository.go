package repository

import "github.com/medvedevse/quote-service/internal/model"

func QuoteRepository() []model.Quote {
	var quotes []model.Quote

	testQuote := model.Quote{
		ID:     "1",
		Quote:  "Life is simple, but we insist on making it complicated.",
		Author: "Confucius",
	}

	quotes = append(quotes, testQuote)
	return quotes
}
