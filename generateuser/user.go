package generateuser

import (
	"github.com/jaswdr/faker/v2"
)

// User represents a user with basic information fields.
// All fields are exported and include JSON tags for serialization.
type User struct {
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
