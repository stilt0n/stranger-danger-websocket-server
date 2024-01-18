package main

import (
	"log"
	"stranger-danger/db"

	// Need to include postgres driver for database to work
	// this isn't used so we need to alias it w/ _
	_ "github.com/lib/pq"
)

func main() {
	database, err := db.NewDatabase()
	if err != nil {
		log.Fatalf("Could not connect to database: %s", err)
		return
	}
	defer database.Close()
	log.Printf("Successfully connected to the database!")
}
