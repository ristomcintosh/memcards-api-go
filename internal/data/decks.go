package data

import (
	"time"

	"memcards.ristomcintosh.com/internal/validator"
)

type Deck struct {
	ID         uint        `json:"id"`
	Name       string      `json:"name"`
	Flashcards []Flashcard `json:"flashcards"`
	CreatedAt  time.Time   `json:"-"`
	UpdatedAt  time.Time   `json:"-"`
}

func ValidateDeck(v *validator.Validator, deck *Deck) {

	v.Check(deck.Name != "", "name", "Missing field: name is required")
	v.Check(len(deck.Name) >= 3, "name", "name should be at least 3 characters long")
}
