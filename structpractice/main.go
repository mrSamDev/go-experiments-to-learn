package main

import (
	"fmt"

	"example.com/app/note"
	"example.com/app/utils"
)

func main() {
	title, content := utils.ReadNote()

	newNote, error := note.New(title, content)

	if error != nil {
		fmt.Println(error)
	}

	newNote.Display()

	error = newNote.Save()

	if error != nil {
		fmt.Println("Saving the nodte failed")
		return
	}

	fmt.Println("Saving the note completed")

}
