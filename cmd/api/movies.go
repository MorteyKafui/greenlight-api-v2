package main

import (
	"fmt"
	"net/http"
	"time"

	"greenlight-api-v2.kafui.net/internal/data"
	"greenlight-api-v2.kafui.net/internal/data/validator"
)

func (app *application) createMovieHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Title   string       `json:"title"`
		Year    int32        `json:"year"`
		Runtime data.Runtime `json:"runtime"`
		Genres  []string     `json:"genres"`
	}

	err := app.readJSON(w, r, &input)

	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	movie := &data.Movie{
		Title:   input.Title,
		Year:    input.Year,
		Runtime: input.Runtime,
		Genres:  input.Genres,
	}

	v := validator.New()

	if data.ValidateMovie(v, movie); !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	fmt.Fprintf(w, "%+v\n", input)
}

func (app *application) showMovieHandler(w http.ResponseWriter, r *http.Request) {

	id, err := app.readIDParam(r)

	if err != nil || id < 1 {
		app.notFoundResponse(w, r)
		return
	}

	movie := data.Movie{
		ID:       id,
		Title:    "A Working Man",
		Runtime:  102,
		Genres:   []string{"action", "drama"},
		Version:  1,
		CreateAt: time.Now(),
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"data": movie}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
