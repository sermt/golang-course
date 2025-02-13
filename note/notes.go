package note

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"
)

type Note struct {
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}

func New() (*Note, error) {
	title, err := getUserInput("Please enter the title:")
	if err != nil {
		return nil, errors.New("invalid title")
	}
	content, err := getUserInput("Please enter the content:")
	if err != nil {
		return nil, errors.New("invalid content")
	}

	fmt.Println("Creating new note...")
	time.Sleep(1 * time.Second)
	fmt.Println("Done.")

	return &Note{
		Title:     title,
		Content:   content,
		CreatedAt: time.Now(),
	}, nil
}

func (note *Note) Save() error {
	fileName := strings.ReplaceAll(strings.ToLower(note.Title), " ", "_") + ".json"
	jsonData, err := json.Marshal(note)
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

func (n *Note) ShowNote() {
	fmt.Printf("Title: %s\nContent: %s\nCreated at: %s\n",
		n.Title, n.Content, n.CreatedAt.Format(time.RFC3339))
}
