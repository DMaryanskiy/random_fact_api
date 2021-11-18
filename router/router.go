package router

import (
	"database/sql"

	"github.com/DMaryanskiy/random_fact_api/facts"
	"github.com/go-chi/chi/v5"
)

func Setup(db *sql.DB) *chi.Mux {
	// setting routes
	r := chi.NewRouter()
	r.Route("/fact", func(r chi.Router) {
		r.Get("/", facts.FactGet(db))
		r.Post("/", facts.FactListPost(db))

		r.Get("/{fact_id}", facts.FactRetreive(db))
		r.Put("/{fact_id}", facts.FactPut(db))
	})

	return r
}
