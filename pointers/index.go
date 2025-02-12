package main

import "fmt"

func main() {
	age := 30
	useAge := &age
	fmt.Println(useAge)
	fmt.Println(getAdultAge(useAge))
}

func getAdultAge(age *int) int {
	return *age - 18
}
