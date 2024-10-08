package data

import (
	"errors"
	"time"

	"gorm.io/gorm"
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

type DeckModel struct {
	DB *gorm.DB
}

func (d DeckModel) Create(deck *Deck) error {
	err := d.DB.Create(&deck).Error

	if err != nil {
		return err
	}

	return nil
}

func (d DeckModel) Update(deck *Deck) error {

	tx := d.DB.Model(&deck).Updates(&deck)

	err := tx.Error
	if err != nil {
		return err
	}

	if tx.RowsAffected == 0 {
		return ErrNoRecord
	}

	return nil
}

func (d DeckModel) GetAll() ([]Deck, error) {
	var decks []Deck
	result := d.DB.Model(&Deck{}).Preload("Flashcards").Find(&decks)
	return decks, result.Error
}

func (d DeckModel) GetByID(id uint) (*Deck, error) {
	var deck Deck

	err := d.DB.First(&deck, id).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrNoRecord
		} else {
			return nil, err
		}
	}

	return &deck, nil
}
