package main

import (
	"fmt"
	"time"
)

type User struct {
	FirstName string
	LastName  string
	Birthdate string
	CreatedAt time.Time
}

func (user *User) outputUserDetail() {
	fmt.Printf("First Name: %s\n", user.FirstName)
	fmt.Printf("Last Name: %s\n", user.LastName)
	fmt.Printf("Birthdate: %s\n", user.Birthdate)
	fmt.Printf("Created At: %s\n", user.CreatedAt)
}

func (user *User) clearUsername() {

	user.FirstName = ""
	user.LastName = ""
}

func NewUser(firstName, lastName, birthdate string) *User {
	return &User{
		FirstName: firstName,
		LastName:  lastName,
		Birthdate: birthdate,
		CreatedAt: time.Now(),
	}

}

func main() {
	// userFirstName := getUserData("Please enter your first name: ")
	// userLastName := getUserData("Please enter your last name: ")
	// userBirthdate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

	user := NewUser("Burhanudin", "Rabbani", "11/14/2002")

	user.outputUserDetail()
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scan(&value)
	return value
}
