package data

import (
	"time"

	"memcards.ristomcintosh.com/internal/validator"
)

type Flashcard struct {
	ID        uint      `json:"id"`
	Front     string    `json:"front"`
	Back      string    `json:"back"`
	DeckID    uint      `json:"deckId"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}

func ValidateFlashcard(v *validator.Validator, flashcard *Flashcard) {

	v.Check(flashcard.Front != "", "front", "Missing field: front is required")
	v.Check(len(flashcard.Front) >= 1, "front", "front should have at least 1 character")
	v.Check(flashcard.Back != "", "back", "Missing field: back is required")
	v.Check(len(flashcard.Back) >= 1, "back", "back should have at least 1 character")
}
