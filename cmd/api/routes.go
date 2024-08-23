package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

func (app *application) routes() http.Handler {
	r := mux.NewRouter()

	r.Use(app.requestLogger)

	r.HandleFunc("/decks", app.GetDecks).Methods(http.MethodGet)
	r.HandleFunc("/decks/{deckId}", app.GetDeck).Methods(http.MethodGet)
	r.HandleFunc("/decks", app.CreateDeck).Methods(http.MethodPost)
	r.HandleFunc("/decks/{deckId}", app.UpdateDeck).Methods(http.MethodPut)
	r.HandleFunc("/decks/{deckId}/flashcards", app.CreateFlashcard).Methods(http.MethodPost)

	return r
}
