// all struct models
package facts

// struct for a fact entity from database
type Fact struct {
	Id          int      `json:"id"`
	Title       string   `json:"title"`
	Description string   `json:"description"`
	Links       []string `json:"links,omitempty"`
}

// struct for a slice of facts which will be added to database
type FactsList struct {
	Facts []Fact `json:"facts"`
}

// struct for a slice of id's of entities added to database
type IDs struct {
	Ids []int `json:"ids"`
}
