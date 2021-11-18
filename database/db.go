package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func Setup() *sql.DB {
	// load environmental variables
	err_env := godotenv.Load()
	if err_env != nil {
		log.Fatal(err_env)
	}

	// connecting to database
	connstr := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("HOST"),
		os.Getenv("PORT"),
		os.Getenv("USER"),
		os.Getenv("PASSWORD"),
		os.Getenv("DBNAME"),
	)
	if connstr == "" {
		log.Fatal("there is no db address")
	}

	db, err := sql.Open("postgres", connstr)
	if err != nil {
		log.Fatal(err)
	}
	return db
}
