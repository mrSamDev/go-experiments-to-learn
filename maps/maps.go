package main

import "fmt"

type Name map[string]int

func (Name) Output() {
	fmt.Println("Output")
}

func main() {
	// maps
	// key-value pair
	// key must be unique
	// value can be duplicated
	// map[key]value
	name := Name{
		"first":  1,
		"second": 10,
	}

	name.Output()

	fmt.Println(name["first"])

	for key, value := range name {
		fmt.Println(value, key)
	}

}
