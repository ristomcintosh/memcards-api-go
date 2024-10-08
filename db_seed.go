package main

import (
	"log"
	"slices"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"memcards.ristomcintosh.com/internal/data"
)

func main() {
	db, err := gorm.Open(sqlite.Open("test.db?_foreign_keys=on"), &gorm.Config{
		TranslateError: true,
	})

	if err != nil {
		log.Fatal("Failed to connect to database: ", err)
	}

	err = seedDB(db)
	if err != nil {
		log.Fatal("Failed to seed the database: ", err)
	}

	log.Println("Database seeded successfully.")
}

func seedDB(db *gorm.DB) error {
	db.Exec("DROP TABLE decks")
	db.Exec("DROP TABLE flashcards")

	err := db.AutoMigrate(&data.Deck{}, &data.Flashcard{})
	if err != nil {
		return err
	}

	decks := []data.Deck{
		{Name: "World Capitals"},
		{Name: "Basic Portuguese"},
	}

	err = db.Create(&decks).Error

	if err != nil {
		return err
	}

	worldCapitals := decks[0]
	portugueseBasic := decks[1]

	worldCapitalsCards := []data.Flashcard{
		{Front: "France", Back: "Paris", DeckID: worldCapitals.ID},
		{Front: "Japan", Back: "Tokyo", DeckID: worldCapitals.ID},
		{Front: "Italy", Back: "Rome", DeckID: worldCapitals.ID},
		{Front: "Brazil", Back: "Brasilia", DeckID: worldCapitals.ID},
		{Front: "Canada", Back: "Ottawa", DeckID: worldCapitals.ID},
	}

	portugueseBasicCards := []data.Flashcard{
		{Front: "Hello", Back: "Olá", DeckID: portugueseBasic.ID},
		{Front: "Thank you", Back: "Obrigado", DeckID: portugueseBasic.ID},
		{Front: "Yes", Back: "Sim", DeckID: portugueseBasic.ID},
		{Front: "No", Back: "Não", DeckID: portugueseBasic.ID},
		{Front: "Goodbye", Back: "Adeus", DeckID: portugueseBasic.ID},
	}

	err = db.Create(slices.Concat(worldCapitalsCards, portugueseBasicCards)).Error

	if err != nil {
		return err
	}

	return nil
}
