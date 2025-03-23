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

func printSomething(data any) {
	//any types can be checked
	typedValue, isTrue := data.(int)

	if !isTrue {
		fmt.Println("This is a string", typedValue)
		return
	}

	switch data.(type) {
	case int:
		fmt.Println("This is an integer")
	case string:
		fmt.Println("This is a string")
	case float64:
		fmt.Println("This is a float")
	case bool:
		fmt.Println("This is a boolean")

	}

	fmt.Println(data)

}

func main() {

	printSomething(1)
	printSomething("string")
	printSomething(1.0)
	printSomething(true)

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
