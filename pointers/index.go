package main

import "fmt"

func main() {
	age := 30
	useAge := &age
	fmt.Println(*useAge)
	getAdultAge(useAge)
	fmt.Println(age)
}

func getAdultAge(age *int) {
	*age = *age - 18
}
