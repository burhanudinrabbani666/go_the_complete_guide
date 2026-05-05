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

func main() {
	userFirstName := getUserData("Please enter your first name: ")
	userLastName := getUserData("Please enter your last name: ")
	userBirthdate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

	user := User{
		FirstName: userFirstName,
		LastName:  userLastName,
		Birthdate: userBirthdate,
		CreatedAt: time.Now(),
	}

	outputUserDetail(user)
}

func outputUserDetail(user User) {

	fmt.Println("First Name: ", user.FirstName)
	fmt.Println("Last Name: ", user.LastName)
	fmt.Println("Birthdate: ", user.Birthdate)
	fmt.Println("Created At: ", user.CreatedAt)
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scan(&value)
	return value
}
