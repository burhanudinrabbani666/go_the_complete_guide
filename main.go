package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"go_the_complete_guide/main/note"
	"go_the_complete_guide/main/todo"
)

type Saver interface {
	Save() error
}

type Outputable interface {
	Saver
	Display()
}

func main() {

	title, description := getNoteData()
	todoText := getUserInput("Todo text: ")

	userTodo, todoError := todo.New(todoText)
	if todoError != nil {
		fmt.Println(todoError)
		return
	}

	errTodoSave := outputData(userTodo)
	if errTodoSave != nil {
		return
	}

	userNote, err := note.New(title, description)
	if err != nil {
		fmt.Println(err)
		return
	}

	errorSave := outputData(userNote)
	if errorSave != nil {
		return
	}

}

func outputData(data Outputable) error {
	data.Display()
	return saveData(data)
}

func saveData(data Saver) error {
	err := data.Save()

	if err != nil {
		fmt.Println("Failed to sava Data")
		return err
	}

	fmt.Println("Saving the note successed")
	return nil
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
