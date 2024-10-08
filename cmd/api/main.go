package main

import (
	"log"
	"net/http"
	"os"

	"memcards.ristomcintosh.com/internal/data"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type application struct {
	infoLog  *log.Logger
	errorLog *log.Logger
	models   data.Models
}

func main() {
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	db, err := dbSetup()

	if err != nil {
		errorLog.Fatal(err)
	}

	app := application{
		infoLog:  infoLog,
		errorLog: errorLog,
		models:   data.NewModels(db),
	}

	srv := &http.Server{
		Addr:     ":5757",
		ErrorLog: errorLog,
		Handler:  app.routes(),
	}

	infoLog.Println("listening on port 5757")
	log.Fatal(srv.ListenAndServe())
}

var decks = []data.Deck{
	{Name: "World Capitals"},
	{Name: "Basic Portuguese"},
}

func dbSetup() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// TODO move to a "seed" script
	db.Exec("DROP TABLE decks")
	db.Exec("DROP TABLE flashcards")
	db.AutoMigrate(data.Deck{}, data.Flashcard{})

	db.Create(&decks)

	worldCapitals := decks[0]
	worldCapitalsCards := []data.Flashcard{
		{Front: "France", Back: "Paris", DeckID: worldCapitals.ID},
		{Front: "Japan", Back: "Tokyo", DeckID: worldCapitals.ID},
		{Front: "Italy", Back: "Rome", DeckID: worldCapitals.ID},
		{Front: "Brazil", Back: "Brasilia", DeckID: worldCapitals.ID},
		{Front: "Canada", Back: "Ottawa", DeckID: worldCapitals.ID},
	}
	db.Create(worldCapitalsCards)

	portugueseBasic := decks[1]
	portugueseBasicCards := []data.Flashcard{
		{Front: "Hello", Back: "Olá", DeckID: portugueseBasic.ID},
		{Front: "Thank you", Back: "Obrigado", DeckID: portugueseBasic.ID},
		{Front: "Yes", Back: "Sim", DeckID: portugueseBasic.ID},
		{Front: "No", Back: "Não", DeckID: portugueseBasic.ID},
		{Front: "Goodbye", Back: "Adeus", DeckID: portugueseBasic.ID},
	}

	db.Create(portugueseBasicCards)

	return db, nil
}
