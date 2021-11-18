// function to count a total amount of facts
package facts

import (
	"database/sql"
	"net/http"
)

// function for getting a total amount of facts
func countRows(db *sql.DB, w http.ResponseWriter) int {
	// getting a total count of rows
	query := `SELECT count(*) FROM public."facts";`
	rows_count, err := db.Query(query)
	checkError(w, err)

	// setting a count to an integer variable
	count := 0
	for rows_count.Next() {
		err = rows_count.Scan(&count)
		checkError(w, err)
	}
	return count
}
