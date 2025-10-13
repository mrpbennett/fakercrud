package main

import (
	"encoding/json"
	"fmt"

	"github.com/google/uuid"
	"github.com/jaswdr/faker/v2"
)

// User represents a user with basic information fields.
// All fields are exported and include JSON tags for serialization.
type User struct {
	UUID     string `json:"uuid"`
	FullName string `json:"full_name"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Address  string `json:"address"`
}

// GenerateUser creates a new User with randomly generated fake data.
// It uses the faker library to generate realistic data for each field.
// Returns a pointer to the newly created User.
func GenerateUser() *User {
	f := faker.New()
	return &User{
		UUID:     uuid.New().String(),
		FullName: f.Person().Name(),
		Email:    f.Internet().CompanyEmail(),
		Phone:    f.Phone().Number(),
		Address:  f.Address().City(),
	}
}

// GenerateUsers creates a slice of User pointers with fake data.
// The count parameter specifies how many users to generate.
// Returns a slice containing the generated users.
func GenerateUsers(count int) []*User {
	users := make([]*User, 0, count)
	for range count {
		users = append(users, GenerateUser())
	}
	return users
}

// main is the entry point that generates fake users and prints them as JSON.
// It panics if the JSON marshalling step fails, keeping the example concise.
func main() {
	users := GenerateUsers(5)
	jsonData, err := json.MarshalIndent(users, "", " ")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(jsonData))
}
