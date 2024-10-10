package main

import (
	"errors"
	"net/http"
	"strings"

	"memcards.ristomcintosh.com/internal/data"
	"memcards.ristomcintosh.com/internal/validator"
)

func (app *application) GetDecks(w http.ResponseWriter, r *http.Request) {
	decks, err := app.models.Deck.GetAll()

	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"decks": decks})

	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

}

func (app *application) GetDeck(w http.ResponseWriter, r *http.Request) {
	deckId, err := app.readIDParam(r, deckIdParam)

	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	deck, err := app.models.Deck.GetByID(uint(deckId))

	if err != nil {
		if errors.Is(err, data.ErrNoRecord) {
			app.notFoundResponse(w, r)
		} else {
			app.serverErrorResponse(w, r, err)
		}
		return
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"deck": deck})

	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}
}

func (app *application) CreateDeck(w http.ResponseWriter, r *http.Request) {

	var input struct {
		Name string `json:"name"`
	}

	err := app.readJSON(w, r, &input)

	if err != nil {
		app.errorResponse(w, r, http.StatusBadRequest, err.Error())
		return
	}

	deck := &data.Deck{
		Name: strings.TrimSpace(input.Name),
	}

	v := validator.New()

	if data.ValidateDeck(v, deck); !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	err = app.models.Deck.Create(deck)

	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	err = app.writeJSON(w, http.StatusCreated, envelope{"deck": deck})

	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}
}

func (app *application) UpdateDeck(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Name string `json:"name"`
	}

	deckId, err := app.readIDParam(r, deckIdParam)

	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	err = app.readJSON(w, r, &input)

	if err != nil {
		app.errorResponse(w, r, http.StatusBadRequest, err.Error())
		return
	}

	deck := &data.Deck{
		Name: strings.TrimSpace(input.Name),
		ID:   uint(deckId),
	}

	v := validator.New()

	if data.ValidateDeck(v, deck); !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	err = app.models.Deck.Update(deck)

	if err != nil {
		if errors.Is(err, data.ErrNoRecord) {
			app.notFoundResponse(w, r)
		} else {
			app.serverErrorResponse(w, r, err)
		}
		return
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"deck": deck})

	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}
}

func (app *application) DeleteDeck(w http.ResponseWriter, r *http.Request) {
	deckId, err := app.readIDParam(r, deckIdParam)

	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	err = app.models.Deck.Delete(uint(deckId))

	if err != nil {
		if errors.Is(err, data.ErrNoRecord) {
			app.notFoundResponse(w, r)
		} else {
			app.serverErrorResponse(w, r, err)
		}
		return
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"message": "deck successfully deleted"})

	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}
}

func (app *application) CreateFlashcard(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Front string `json:"front"`
		Back  string `json:"back"`
	}

	err := app.readJSON(w, r, &input)

	if err != nil {
		app.errorResponse(w, r, http.StatusBadRequest, err.Error())
		return
	}

	deckId, err := app.readIDParam(r, deckIdParam)

	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	flashcard := &data.Flashcard{
		Front:  input.Front,
		Back:   input.Back,
		DeckID: uint(deckId),
	}

	v := validator.New()
	if data.ValidateFlashcard(v, flashcard); !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	err = app.models.Flashcard.Create(flashcard)

	if err != nil {
		if errors.Is(err, data.ErrNoRecord) {
			app.notFoundResponse(w, r)
		} else {
			app.serverErrorResponse(w, r, err)
		}
		return
	}

	err = app.writeJSON(w, http.StatusCreated, envelope{"flashcard": flashcard})

	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}
}

func (app *application) UpdateFlashcard(w http.ResponseWriter, r *http.Request) {
	deckId, err := app.readIDParam(r, deckIdParam)

	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	flashcardId, err := app.readIDParam(r, flashcardIdParam)

	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	var input struct {
		Front string `json:"front"`
		Back  string `json:"back"`
	}

	err = app.readJSON(w, r, &input)

	if err != nil {
		app.errorResponse(w, r, http.StatusBadRequest, err.Error())
		return
	}

	flashcard := &data.Flashcard{
		ID:     uint(flashcardId),
		Front:  input.Front,
		Back:   input.Back,
		DeckID: uint(deckId),
	}

	v := validator.New()
	if data.ValidateFlashcard(v, flashcard); !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	err = app.models.Flashcard.Update(flashcard)

	if err != nil {
		if errors.Is(err, data.ErrNoRecord) {
			app.notFoundResponse(w, r)
		} else {
			app.serverErrorResponse(w, r, err)
		}
		return
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"flashcard": flashcard})

	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}
}

func (app *application) DeleteFlashcard(w http.ResponseWriter, r *http.Request) {
	flashcardId, err := app.readIDParam(r, flashcardIdParam)

	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	err = app.models.Flashcard.Delete(uint(flashcardId))

	if err != nil {
		if errors.Is(err, data.ErrNoRecord) {
			app.notFoundResponse(w, r)
		} else {
			app.serverErrorResponse(w, r, err)
		}
		return
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"message": "flashcard successfully deleted"})

	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}
}
