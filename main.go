package main

import (
	"fmt"

	"examples.com/go-project/note"
)

func main() {
	note, err := note.New()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	note.ShowNote()

	if err := note.Save(); err != nil {
		fmt.Println("Error saving note:", err)
	} else {
		fmt.Println("Note saved successfully!")
	}
}
