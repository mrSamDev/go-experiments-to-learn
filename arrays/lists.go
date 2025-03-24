package main

import (
	"fmt"
)

func main() {

	products := [4]string{"apple", "banana", "orange", "grape"}

	picked := [1]string{"apple"}

	// passed by reference
	someProducts := products[1:3]

	fmt.Println(someProducts)

	fmt.Println(len(picked))

	fmt.Println(products)
	dynamicArray()
	exercise()
}

type Product struct {
	Name  string
	Price float64
}

func dynamicArray() {
	products := []string{}
	products = append(products, "apple", "banana", "orange", "grape")
	fmt.Println(products, "[dynamicArray]")
}

func exercise() {
	hobbies := [3]string{"reading", "running", "coding"}
	fmt.Println(hobbies, "[hobbies]")
	fmt.Println(hobbies[0], "[First Hobby]")
	secondAndThirdHobbies := hobbies[1:]
	fmt.Println(hobbies[0], "[First Hobby]")
	fmt.Println(secondAndThirdHobbies, "[second and third hobbies]")
	firstAndSecondHobbies := hobbies[:2]
	fmt.Println(firstAndSecondHobbies, "[first and second hobbies]")

	productList := []Product{
		{Name: "apple"},
		{Name: "banana", Price: 0.99},
		{Name: "orange", Price: 1.49},
	}

	fmt.Println(productList, "[productList]")

	for index, value := range productList {
		fmt.Println(value.Name, index)
	}

}
