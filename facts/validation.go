// validation functions
package facts

import (
	"database/sql"
	"net/http"
)

// validation for POST method
func validatePost(fact Fact) (bool, string) {
	// checking if title and description fields aren't empty
	if fact.Title == "" || fact.Description == "" {
		return false, `fields "title" and "description" must be filled`
	}

	// checking if title and links aren't too long
	title := []rune(fact.Title)
	if len(title) > 255 {
		return false, `field "title" is too long`
	}

	for _, elem := range fact.Links {
		description := []rune(elem)
		if len(description) > 255 {
			return false, "one of the links is too long"
		}
	}

	return true, ""
}

// validation for PUT method
func validatePut(fact_id int, fact Fact, db *sql.DB, w http.ResponseWriter) (bool, string) {
	if fact_id != fact.Id {
		return false, "given id doesn't equal to id in url"
	}

	count := countRows(db, w)
	if fact_id < 1 || fact_id > count {
		return false, "there is no fact with given id"
	}

	if ok, msg := validatePost(fact); !ok {
		return false, msg
	}

	return true, ""
}
