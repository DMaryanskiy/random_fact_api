// PUT method
package facts

import (
	"database/sql"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/lib/pq"
)

// handler for updating a fact with given id
func FactPut(db *sql.DB) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// setting headers
		w.Header().Set("Content-Type", "application/json")

		// initializing temporary and url parameter variables
		fact := Fact{}
		fact_id, err := strconv.Atoi(chi.URLParam(r, "fact_id"))
		checkError(w, err)

		// reading request body
		body, err := ioutil.ReadAll(r.Body)
		checkError(w, err)

		// unmarshalling JSON to a struct
		err = json.Unmarshal(body, &fact)
		checkError(w, err)

		// validating given data
		if ok, msg := validatePut(fact_id, fact, db, w); !ok {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(msg))
			return
		}

		// updating a fact with given id
		query := `UPDATE public."facts" SET title=$1, description=$2, links=$3 WHERE id=$4`
		upd_fact, err := db.Exec(query, fact.Title, fact.Description, pq.Array(fact.Links), strconv.Itoa(fact_id))
		_ = upd_fact
		checkError(w, err)

		w.Write([]byte("{}"))
	})
}
