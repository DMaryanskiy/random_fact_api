package main

import (
	"log"
	"net/http"

	"github.com/DMaryanskiy/random_fact_api/database"
	"github.com/DMaryanskiy/random_fact_api/router"
)

func main() {
	// setting up a db
	db := database.Setup()
	defer db.Close()

	// setting up a router
	r := router.Setup(db)
	log.Fatal(http.ListenAndServe(":3000", r))
}
