package main

import (
	"errors"
	"net/http"
	"strings"

	"memcards.ristomcintosh.com/internal/data"
	"memcards.ristomcintosh.com/internal/validator"
)

func (app *application) GetDecks(w http.ResponseWriter, r *http.Request) {
	decks, err := app.db.GetAllDecks()

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
	deckId, err := app.readIDParam(r)

	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	deck, err := app.db.GetDeckByID(uint(deckId))

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

	newDeck, err := app.db.CreateDeck(deck.Name)

	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	err = app.writeJSON(w, http.StatusCreated, envelope{"deck": newDeck})

	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}
}

func (app *application) UpdateDeck(w http.ResponseWriter, r *http.Request) {
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

	deckId, err := app.readIDParam(r)

	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	updatedDeck, err := app.db.UpdateDeck(uint(deckId), deck.Name)

	if err != nil {
		if errors.Is(err, data.ErrNoRecord) {
			app.notFoundResponse(w, r)
		} else {
			app.serverErrorResponse(w, r, err)
		}
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"deck": updatedDeck})

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

	deckId, err := app.readIDParam(r)

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

	newFlashcard, err := app.db.CreateFlashcard(flashcard.DeckID, flashcard.Front, flashcard.Back)

	if err != nil {
		if errors.Is(err, data.ErrNoRecord) {
			app.notFoundResponse(w, r)
		} else {
			app.serverErrorResponse(w, r, err)
		}
		return
	}

	err = app.writeJSON(w, http.StatusCreated, envelope{"flashcard": newFlashcard})

	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}
}
