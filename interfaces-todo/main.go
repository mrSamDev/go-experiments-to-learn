package main

import (
	"fmt"

	"example.com/app/note"
	"example.com/app/todo"
	"example.com/app/utils"
)

type ouputtable interface {
	Display()
	Save() error
}

func main() {
	title := utils.ReadTodo()

	newtodo, error := todo.New(title)

	if error != nil {
		fmt.Println(error)
	}

	error = saveDataAndDisplay(newtodo)

	if error != nil {
		fmt.Println(error)
		return
	}

	title, content := utils.ReadNote()

	newNote, error := note.New(title, content)

	if error != nil {
		fmt.Println(error)
	}

	error = saveDataAndDisplay(newNote)

	if error != nil {
		fmt.Println(error)
	}

}

func saveDataAndDisplay(data ouputtable) error {

	data.Display()

	error := data.Save()

	if error != nil {
		fmt.Println("Saving failed")
		return nil
	}

	fmt.Println("Saving completed")

	return nil
}
