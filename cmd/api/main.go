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

func dbSetup() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("test.db?_foreign_keys=on"), &gorm.Config{
		TranslateError: true,
	})

	if err != nil {
		return nil, err
	}

	return db, nil
}
