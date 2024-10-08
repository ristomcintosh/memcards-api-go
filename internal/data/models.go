package data

import "gorm.io/gorm"

type Models struct {
	Deck      DeckModel
	Flashcard FlashcardModel
}

func NewModels(db *gorm.DB) Models {
	return Models{
		Deck:      DeckModel{DB: db},
		Flashcard: FlashcardModel{DB: db},
	}
}
