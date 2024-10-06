package main

import (
	"errors"
	"net/http"

	"memcards.ristomcintosh.com/internal/data"
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

	// if req.Name == "" {
	// 	w.WriteHeader(http.StatusBadRequest)
	// 	w.Write([]byte("name is required"))
	// 	return
	// }

	newDeck, _ := app.db.CreateDeck(input.Name)

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

	if input.Name == "" {
		app.errorResponse(w, r, http.StatusBadRequest, "name is required")
		return
	}

	deckId, err := app.readIDParam(r)

	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	deck, err := app.db.UpdateDeck(uint(deckId), input.Name)

	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	if deck == nil {
		app.notFoundResponse(w, r)
		return
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"deck": deck})

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

	newFlashcard, _ := app.db.CreateFlashcard(uint(deckId), input.Front, input.Back)

	// if deck == nil {
	// 	w.WriteHeader(http.StatusNotFound)
	// 	message := fmt.Sprintf("Deck: %v not found", deckId)
	// 	json.NewEncoder(w).Encode(APIResponse{
	// 		Message: message,
	// 	})
	// 	return
	// }

	err = app.writeJSON(w, http.StatusCreated, envelope{"flashcard": newFlashcard})

	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}
}
