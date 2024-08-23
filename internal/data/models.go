package data

import (
	"time"
)

type Deck struct {
	ID         uint        `json:"id"`
	Name       string      `json:"name"`
	Flashcards []Flashcard `json:"flashcards"`
	CreatedAt  time.Time   `json:"-"`
	UpdatedAt  time.Time   `json:"-"`
}

type Flashcard struct {
	ID        uint      `json:"id"`
	Front     string    `json:"front"`
	Back      string    `json:"back"`
	DeckID    uint      `json:"deckId"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
}
