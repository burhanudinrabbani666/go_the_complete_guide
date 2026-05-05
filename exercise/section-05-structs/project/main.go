package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"go_the_complete_guide/main/note"
)

func main() {

	title, description := getNoteData()
	userNote, err := note.New(title, description)

	if err != nil {
		fmt.Println(err)
		return
	}

	userNote.Display()
	errorSave := userNote.Save()
	if errorSave != nil {
		fmt.Println("Saving the note Failed")
		return
	}

	fmt.Println("Saving the note Succeeded")
}

func getNoteData() (string, string) {
	title := getUserInput("Note title: ")
	description := getUserInput("Note Description: ")

	return title, description

}

func getUserInput(prompt string) string {
	fmt.Print(prompt)

	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')

	if err != nil {
		return ""
	}

	text = strings.Trim(text, "\n")
	return text
}
