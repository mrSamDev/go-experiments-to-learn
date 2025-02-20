package main

import (
	"fmt"

	"example.com/app/note"
	"example.com/app/utils"
)

func main() {
	title, content, error := utils.ReadNote()
	if error != nil {
		fmt.Println(error)
	}

	newNote, error := note.New(title, content)

	if error != nil {
		fmt.Println(error)
	}

	newNote.Display()

}
