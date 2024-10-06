package main

import (
	"fmt"
	"net/http"
	"runtime/debug"
)

func (app *application) logError(r *http.Request, err error) {
	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
	app.errorLog.Output(2, trace)
}

func (app *application) errorResponse(w http.ResponseWriter, r *http.Request, status int, message any) {
	errMsg := envelope{"error": message}

	err := app.writeJSON(w, status, errMsg)

	if err != nil {
		app.logError(r, err)
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func (app *application) serverErrorResponse(w http.ResponseWriter, r *http.Request, err error) {
	app.logError(r, err)
	message := "the server encountered a problem and could not process your request"
	app.errorResponse(w, r, http.StatusInternalServerError, message)
}

func (app *application) notFoundResponse(w http.ResponseWriter, r *http.Request) {
	message := "the requested resource could not be found"
	app.errorResponse(w, r, http.StatusNotFound, message)
}
