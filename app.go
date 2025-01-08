package main

import (
	"fmt"
	"math"
)

func investmentCalculator() {
	const inflation = 1.4
	var investmentAmount, years, expectedReturnRate float64

	fmt.Print("Enter the investment amount: ")
	fmt.Scan(&investmentAmount)

	fmt.Print("Enter the number of years: ")
	fmt.Scan(&years)

	fmt.Print("Enter the expected return rate: ")

	fmt.Scan(&expectedReturnRate)

	future := investmentAmount * math.Pow(1+expectedReturnRate/100, years)

	futureRealValue := future / math.Pow(1+inflation/100, years)

	fmt.Printf("Future %v \n Future Real Value %v", future, futureRealValue)

}

func taxCalculator() {
	const depreciation = 1.4

	income := numberInput("Enter the income: ")
	taxRate := numberInput("Enter the tax rate: ")
	expenses := numberInput("Enter the expenses: ")

	fmt.Println("income", income)
	fmt.Print("Enter the tax rate: ")
	fmt.Scan(&taxRate)
	fmt.Print("Enter the expenses: ")
	fmt.Scan(&expenses)
	ebit := income - expenses
	tax := income * taxRate / 100
	taxRealValue := tax / depreciation
	profit := ebit - tax
	fmt.Println("tax", tax)
	fmt.Println("taxRealValue", taxRealValue)
	fmt.Println("ebit", ebit)
	fmt.Println("profit", profit)
	output("Tax calculator")

}

func main() {
	output("For investment calcaultor press: 1")
	output("For tax calcaultor press: 2")

	var selectedtype int16

	fmt.Scan(&selectedtype)

	switch selectedtype {
	case 1:
		investmentCalculator()
	case 2:
		taxCalculator()
	}
}

func output(text string) {
	fmt.Println(text)
}

func numberInput(text string) (input float64) {
	output(text)
	fmt.Scan(&input)
	return input
}
