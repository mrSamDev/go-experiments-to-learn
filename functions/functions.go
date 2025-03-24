package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3, 4, 5}
	doubrNumber := doubleNumber(&numbers)
	transformNumber(&numbers, double)
	transformNumber(&numbers, func(value int) int {
		return value * 10
	})
	fmt.Println(doubrNumber)
}

func double(number int) int {
	return number * 2
}

func doubleNumber(numbers *[]int) []int {
	var doubleNumbers []int

	for _, value := range *numbers {
		doubleNumbers = append(doubleNumbers, value*2)
	}

	return doubleNumbers

}

type transformfn func(int) int

func transformNumber(numbers *[]int, transform transformfn) []int {

	var doubleNumbers []int

	for _, value := range *numbers {
		doubleNumbers = append(doubleNumbers, transform(value))
	}

	return doubleNumbers
}
