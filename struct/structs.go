package main

import (
	"fmt"

	"example.com/structs/person"
)

func main() {
	p, error := person.New(getUserData("Please enter your first name: "), getUserData("Please enter your last name: "), getUserData("Please enter your birthdate (MM/DD/YYYY): "))

	if error != nil {
		fmt.Println(error)
		return
	}

	p.CapitalizeFirstName()
	p.Show()

}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scanln(&value)
	return value
}

/*
	just need to make the value is in order
		p = person{
			getUserData("Please enter your first name: "),
			getUserData("Please enter your last name: "),
			getUserData("Please enter your birthdate (MM/DD/YYYY): "),
			time.Now(),
		}

		p = person{}
*/
