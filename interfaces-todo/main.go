package main

import (
	"fmt"

	"example.com/app/note"
	"example.com/app/todo"
	"example.com/app/utils"
)

type saver interface {
	Save() error
}

func main() {
	title := utils.ReadTodo()

	newtodo, error := todo.New(title)

	if error != nil {
		fmt.Println(error)
	}

	newtodo.Display()

	error = saveData(newtodo)

	if error != nil {
		fmt.Println(error)
		return
	}

	title, content := utils.ReadNote()

	newNote, error := note.New(title, content)

	if error != nil {
		fmt.Println(error)
	}

	newNote.Display()

	error = saveData(newNote)

	if error != nil {
		fmt.Println(error)
	}

}

func saveData(data saver) error {

	error := data.Save()

	if error != nil {
		fmt.Println("Saving failed")
		return nil
	}

	fmt.Println("Saving completed")

	return nil
}
