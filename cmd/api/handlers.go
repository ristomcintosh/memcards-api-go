package main

import (
	"errors"
	"memcards.ristomcintosh.com/internal/data"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
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
	vars := mux.Vars(r)
	id := vars["deckId"]

	deck, err := app.db.GetDeckByID(id)

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
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("name is required"))
		return
	}

	vars := mux.Vars(r)
	// TODO handle error
	id, _ := strconv.Atoi(vars["deckId"])

	deck, _ := app.db.UpdateDeck(uint(id), input.Name)

	// if deck == nil {
	// 	w.WriteHeader(http.StatusNotFound)
	// 	message := fmt.Sprintf("Deck: %v not found", id)
	// 	json.NewEncoder(w).Encode(APIResponse{
	// 		Message: message,
	// 	})
	// 	return
	// }

	err = app.writeJSON(w, http.StatusOK, envelope{"deck": deck})

	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}
}

func (app *application) CreateFlashcard(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	deckId, _ := strconv.Atoi(vars["deckId"])

	var input struct {
		Front string `json:"front"`
		Back  string `json:"back"`
	}

	err := app.readJSON(w, r, &input)

	if err != nil {
		app.errorResponse(w, r, http.StatusBadRequest, err.Error())
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
