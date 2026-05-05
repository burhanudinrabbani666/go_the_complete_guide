package main

import (
	"errors"
	"fmt"
)

func main() {

	title, description, err := getNoteData()

	if err != nil {
		fmt.Println(err)
		return
	}

}

func getNoteData() (string, string, error) {
	title, errTitle := getUserInput("Note title: ")

	if errTitle != nil {
		fmt.Println(errTitle)
		return "", "", errTitle
	}

	description, errDescription := getUserInput("Note Description: ")
	if errDescription != nil {
		fmt.Println(errTitle)
		return "", "", errDescription
	}

	return title, description, nil

}

func getUserInput(prompt string) (string, error) {
	fmt.Print(prompt)

	var value string
	fmt.Scanln(&value)

	if value == "" {
		return "", errors.New("Input Not Valid")
	}

	return value, nil
}
