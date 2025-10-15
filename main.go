package main

import "github.com/mrpbennett/fakercrud/db"

// main is the entry point that generates fake users and prints them as JSON.
// It panics if the JSON marshalling step fails, keeping the example concise.
func main() {
	// TODO: flush out routes

	// CREATE inital table before user can hit the endpoints...
	db.CreateInitialTable()

	// ...
	db.CreateUsers(5)
}
