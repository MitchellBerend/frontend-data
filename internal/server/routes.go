package server

import (
	"bytes"
	"frontend-data/internal/files"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

var NotFound = []byte("Could not find data")

func (s *Server) RegisterRoutes() http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://api.berendhuysen.nl", "http://localhost"},
		AllowedMethods:   []string{"GET", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	r.Group(func(router chi.Router) {
		router.Use(middleware.SetHeader("Content-Type", "application/json"))
		router.Get("/{data:[a-zA-Z]*-?[a-zA-Z]*}", s.data)
	})

	return r
}

// This is the main function of the "backend". It simply finds the matching file
// for the given parameter and returns it's content with the header
// Content-type set to application/json.
func (s *Server) data(w http.ResponseWriter, r *http.Request) {
	dataParam := chi.URLParam(r, "data")
	if dataParam == "" {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	data := findData(dataParam)

	if bytes.Equal(data, NotFound) {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	_, err := w.Write(data)
	if err != nil {
		log.Printf("Could not write back data in /index: %e", err)
		w.WriteHeader(http.StatusInternalServerError)
	}
}

// A simple switch case to match file content with a given string
func findData(file string) []byte {
	switch file {
	case "index":
		return files.IndexJson
	case "over-engineering":
		return files.OverEngineeringJson
	default:
		return NotFound
	}
}
