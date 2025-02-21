package todo

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"
)

type Todo struct {
	Text string `json:"text"`
}

func New() (*Todo, error) {
	text, err := getUserInput("Please enter the text:")
	if err != nil {
		return nil, errors.New("invalid title")
	}

	fmt.Println("Creating new todo...")
	time.Sleep(1 * time.Second)
	fmt.Println("Done.")

	return &Todo{
		Text: text,
	}, nil
}

func (todo *Todo) Save() error {
	fileName := "todo.json"
	jsonData, err := json.Marshal(todo)
	if err != nil {
		fmt.Printf("Error marshaling JSON: %v\n", err)
		return err
	}

	return os.WriteFile(fileName, jsonData, 0644)
}

func getUserInput(prompt string) (string, error) {
	fmt.Println(prompt)
	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')
	text = strings.TrimSpace(text)

	if err != nil || text == "" {
		return "", errors.New("input cannot be empty")
	}
	return text, nil
}

func (todo *Todo) Display() {
	fmt.Printf("Text: %s\n",
		todo.Text)
}
