// POST handler
package facts

import (
	"database/sql"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/lib/pq"
)

// handler for adding new facts to database
func FactListPost(db *sql.DB) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// setting headers
		w.Header().Set("Content-Type", "application/json")

		// initialization of all temporary variables
		facts := FactsList{}
		ids := IDs{}

		// reading request body
		body, err := ioutil.ReadAll(r.Body)
		checkError(w, err)

		// unmarshalling JSON to a struct
		err = json.Unmarshal(body, &facts)
		checkError(w, err)

		for _, elem := range facts.Facts {
			if ok, msg := validatePost(elem); !ok {
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte(msg))
				return
			}
		}

		// inserting new rows to database and returning ids of new entities
		query := `INSERT INTO public."facts" (title, description, links) VALUES ($1, $2, $3) RETURNING id;`
		for _, elem := range facts.Facts {
			err = db.QueryRow(query, elem.Title, elem.Description, pq.Array(elem.Links)).Scan(&elem.Id)
			checkError(w, err)
			ids.Ids = append(ids.Ids, elem.Id)
		}

		// marshalling new ids
		resp, err := json.Marshal(ids)
		checkError(w, err)

		// writing a response
		w.Write(resp)
	})
}
