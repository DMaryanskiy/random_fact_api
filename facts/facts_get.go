// GET method
package facts

import (
	"database/sql"
	"encoding/json"
	"math/rand"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/lib/pq"
)

// handler for getting random fact
func FactGet(db *sql.DB) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// setting headers
		w.Header().Set("Content-Type", "application/json")

		// initialize temporary variable
		fact := Fact{}

		// getting a random fact
		random_id := rand.Intn(countRows(db, w)) + 1
		query := `SELECT * FROM public."facts" WHERE id = $1;`
		fact_db, err := db.Query(query, random_id)
		checkError(w, err)

		// writing a random fact to a struct
		for fact_db.Next() {
			err = fact_db.Scan(
				&fact.Id,
				&fact.Title,
				&fact.Description,
				pq.Array(&fact.Links),
			)
			checkError(w, err)
		}

		// marshalling struct to JSON
		resp, err := json.Marshal(fact)
		checkError(w, err)

		// writing a response
		w.Write(resp)
	})
}

// handler for retreiving a fact with given id
func FactRetreive(db *sql.DB) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// setting headers
		w.Header().Set("Content-Type", "application/json")

		// initialization of temporary variable
		fact := Fact{}

		// retreiving a fact from database
		fact_id := chi.URLParam(r, "fact_id")
		query := `SELECT * FROM public."facts" WHERE id = $1;`
		fact_db, err := db.Query(query, fact_id)
		checkError(w, err)

		// writing a random fact to a struct
		for fact_db.Next() {
			err = fact_db.Scan(
				&fact.Id,
				&fact.Title,
				&fact.Description,
				pq.Array(&fact.Links),
			)
			checkError(w, err)
		}

		// marshalling struct to JSON
		resp, err := json.Marshal(fact)
		checkError(w, err)

		// writing a response
		w.Write(resp)
	})
}
