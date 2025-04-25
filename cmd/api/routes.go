package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (app *application) routes() http.Handler {
	router := httprouter.New()

	router.NotFound = http.HandlerFunc(app.notFoundResponse)
	router.MethodNotAllowed = http.HandlerFunc(app.methodNotAllowedResponse)

	router.HandlerFunc(http.MethodGet, "/api/v1/healthcheck", app.healthcheckHandler)
	router.HandlerFunc(http.MethodPost, "/api/v1/movies", app.createMovieHandler)
	router.HandlerFunc(http.MethodGet, "/api/v1/movies/:id", app.showMovieHandler)
	router.HandlerFunc(http.MethodPatch, "/api/v1/movies/:id", app.updateMovieHandler)
	router.HandlerFunc(http.MethodDelete, "/api/v1/movies/:id", app.deleteMovieHandler)

	return app.recoverPanic(router)

}
