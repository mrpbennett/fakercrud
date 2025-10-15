package main

import (
	"encoding/json"
	"fmt"

	"github.com/mrpbennett/fakercrud/db"
	"github.com/mrpbennett/fakercrud/generateuser"
)

// main is the entry point that generates fake users and prints them as JSON.
// It panics if the JSON marshalling step fails, keeping the example concise.
func main() {
	users := generateuser.GenerateUsers(1)
	jsonData, err := json.MarshalIndent(users, "", " ")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(jsonData))

	// CREATE table
	db.CreateInitialTable()
}
