package main

import "fmt"

func main() {
	// result := add(1, 2)
	// fmt.Println(result)
	// value := transform(1)(10)
	fmt.Println(factorial(7))
}

// func add[T int | float64](a, b T) T {
// 	return a + b
// }

func factorial(number int) int {

	if number == 0 {
		fmt.Println("exit", number)
		return 1
	}

	return number * factorial(number-1)

}

// func transform(times int) func(int) int {
// 	return func(number int) int {
// 		return number * times
// 	}
// }

// 5 * 4 * 3 * 2 * 1
