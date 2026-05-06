package main

import (
	"fmt"
	"go_the_complete_guide/structs/user"
)

func main() {
	userFirstName := getUserData("Please enter your first name: ")
	userLastName := getUserData("Please enter your last name: ")
	userBirthdate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

	newUser, err := user.New(userFirstName, userLastName, userBirthdate)

	if err != nil {
		fmt.Println(err)
		return
	}
	newUser.OutputUserDetail()
	newAdmin := user.NewAdmin("xxxxxxx", "123455667")
	// newAdmin.ClearUsername()

	fmt.Println(newAdmin)

}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scanln(&value)
	return value
}
