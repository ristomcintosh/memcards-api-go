package data

import (
	"time"
)

type Flashcard struct {
	ID        uint      `json:"id"`
	Front     string    `json:"front"`
	Back      string    `json:"back"`
	DeckID    uint      `json:"deckId"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}
