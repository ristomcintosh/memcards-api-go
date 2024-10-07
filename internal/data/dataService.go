package data

import (
	"errors"

	"gorm.io/gorm"
)

type DataService interface {
	GetAllDecks() ([]Deck, error)
	GetDeckByID(id uint) (*Deck, error)
	CreateDeck(name string) (*Deck, error)
	UpdateDeck(id uint, name string) (*Deck, error)
	CreateFlashcard(deckId uint, front, back string) (*Flashcard, error)
}

type GormOrm struct {
	*gorm.DB
}

func (orm *GormOrm) GetAllDecks() ([]Deck, error) {
	var decks []Deck
	result := orm.Model(&Deck{}).Preload("Flashcards").Find(&decks)
	return decks, result.Error
}

func (orm *GormOrm) GetDeckByID(id uint) (*Deck, error) {
	var deck Deck

	err := orm.First(&deck, id).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, ErrNoRecord
		} else {
			return nil, err
		}
	}

	return &deck, nil
}

func (orm *GormOrm) CreateDeck(name string) (*Deck, error) {
	deck := Deck{Name: name}
	err := orm.Create(&deck).Error

	if err != nil {
		return nil, err
	}

	return &deck, nil
}

func (orm *GormOrm) UpdateDeck(id uint, name string) (*Deck, error) {
	deck := Deck{ID: id}

	tx := orm.Model(&deck).Update("name", name)

	err := tx.Error
	if err != nil {
		return nil, err
	}

	if tx.RowsAffected == 0 {
		return nil, ErrNoRecord
	}

	return &deck, nil
}

func (orm *GormOrm) CreateFlashcard(deckId uint, front, back string) (*Flashcard, error) {

	deck, err := orm.GetDeckByID(deckId)

	if err != nil {
		return nil, err
	}

	flashcard := Flashcard{Front: front, Back: back, DeckID: deck.ID}

	err = orm.Create(&flashcard).Error

	if err != nil {
		return nil, err
	}

	return &flashcard, nil
}
