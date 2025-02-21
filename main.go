package main

import (
	"fmt"

	"examples.com/go-project/note"
	"examples.com/go-project/todo"
)

type saver interface {
	Save() error
}

type displayer interface {
	Display()
}

type outputtable interface {
	saver
	Display()
}

/* type outputtable interface {
	saver
	displayer
} */

func main() {
	note, err := note.New()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	note.ShowNote()

	saveData(note)

	todo, err := todo.New()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	outputData(todo)

	fmt.Println("The larger number is:", printResult(10, 5))

}

func saveData(data saver) error {
	err := data.Save()

	if err != nil {
		fmt.Println("Saving the note failed:", err)
		return err
	}
	fmt.Println("Data saved successfully!")
	return nil
}

func outputData(data outputtable) {
	data.Display()
	saveData(data)
}

func printResult[T int | string](a, b T) T {
	if a > b {
		return a
	}
	return b
}
