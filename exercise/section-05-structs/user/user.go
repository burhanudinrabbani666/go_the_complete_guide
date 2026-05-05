package user

import (
	"errors"
	"fmt"
	"time"
)

type User struct {
	FirstName string
	LastName  string
	Birthdate string
	CreatedAt time.Time
}

type Admin struct {
	Email    string
	Password string
	User
}

func (user *User) OutputUserDetail() {
	fmt.Printf("First Name: %s\n", user.FirstName)
	fmt.Printf("Last Name: %s\n", user.LastName)
	fmt.Printf("Birthdate: %s\n", user.Birthdate)
	fmt.Printf("Created At: %s\n", user.CreatedAt)
}

func (user *User) ClearUsername() {
	user.FirstName = ""
	user.LastName = ""
}

func New(firstName, lastName, birthdate string) (*User, error) {

	if firstName == "" || lastName == "" || birthdate == "" {
		return nil, errors.New("First Name, Last Name, Birthdate is required")
	}

	return &User{
		FirstName: firstName,
		LastName:  lastName,
		Birthdate: birthdate,
		CreatedAt: time.Now(),
	}, nil

}

func NewAdmin(email, password string) Admin {
	return Admin{
		Email:    email,
		Password: password,
		User: User{
			FirstName: "ADMIN",
			LastName:  "ADMIN",
			Birthdate: "---------",
			CreatedAt: time.Now(),
		},
	}
}
