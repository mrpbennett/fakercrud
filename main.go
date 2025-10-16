package main

import (
	"github.com/mrpbennett/fakercrud/db"
)

// main is the entry point that generates fake users and prints them as JSON.
// It panics if the JSON marshalling step fails, keeping the example concise.
func main() {
	// TODO: flush out routes

	// CREATE inital table before user can hit the endpoints...
	err := db.CreateInitialTable()
	if err != nil {
		return
	} else {
		// Populate the user table with 50 users, allowing the user
		// to use the CRUD endpoints.
		db.InitialiseUsers(50)
	}
}
