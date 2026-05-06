package todo

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

type Todo struct {
	Text string `json:"todo"`
}

func (todo Todo) Display() {
	fmt.Println(todo.Text)
}

func New(content string) (*Todo, error) {

	if content == "" {
		return &Todo{}, errors.New("Input Not Valid")
	}

	return &Todo{
		Text: content,
	}, nil

}

func (todo Todo) Save() error {
	fileName := "todo.json"

	bytes, err := json.Marshal(todo)
	if err != nil {
		return err
	}

	return os.WriteFile(fileName, bytes, 0666)
}
