package main

import (
	"encoding/json"
	"io"
	"net/http"

	log "github.com/sirupsen/logrus"
)

type MoviesHttp struct {
	db *MovieDb
}

func CreateMovieHttp(db *MovieDb) *MoviesHttp {
	return &MoviesHttp{db}
}

func (m *MoviesHttp) Handler(rw http.ResponseWriter, req *http.Request) {
	if req.Method == http.MethodGet {
		m.getAll(rw, req)
	} else if req.Method == http.MethodPost {
		m.create(rw, req)
	} else {
		MethodNotAllowed(rw)
	}
}

func (m *MoviesHttp) getAll(rw http.ResponseWriter, req *http.Request) {
	movies, err := m.db.SelectMovies()
	if err != nil {
		log.Errorf("Cannot select movies: %v", err)
		InternalServerError(rw)
		return
	}

	body, mErr := json.Marshal(movies)
	if mErr != nil {
		log.Errorf("Cannot select movies: %v", mErr)
		InternalServerError(rw)
		return
	}

	SuccessString(rw, string(body))
}

func (m *MoviesHttp) create(rw http.ResponseWriter, req *http.Request) {
	body, err := io.ReadAll(req.Body)
	if err != nil {
		log.Errorf("Cannot read body! %v", err)
		InternalServerError(rw)
		return
	}

	var movie Movie
	err = json.Unmarshal(body, &movie)
	if err != nil {
		log.Errorf("Cannot read body! %v", err)
		InternalServerError(rw)
		return
	}

	err = m.db.CreateMovie(movie)
	if err != nil {
		log.Errorf("Cannot create new movie in database! %v", err)
	}

	SuccessString(rw, "Success!")
}
