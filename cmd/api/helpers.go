package main

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type envelope map[string]any

func (app *application) writeJSON(w http.ResponseWriter, status int, data envelope) error {
	js, err := json.Marshal(data)
	if err != nil {
		return err
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	w.Write(js)

	return nil
}

func (app *application) readJSON(w http.ResponseWriter, r *http.Request, dst any) error {
	err := json.NewDecoder(r.Body).Decode(dst)
	if err != nil {
		return err
	}

	return nil
}

func (app *application) readIDParam(r *http.Request, paramName routeParam) (int, error) {
	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars[string(paramName)])

	if err != nil {
		return 0, err
	}

	return id, nil
}
