package main

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

const fileName = "response.text"

func main() {
	revenue, rerror := getUserInput("Revenue: ")
	expenses, rexpense := getUserInput("Expenses: ")
	taxRate, rtax := getUserInput("Tax Rate: ")

	if rerror != nil || rexpense != nil || rtax != nil {
		return
	}

	ebt, profit, ratio := calculateFinancials(revenue, expenses, taxRate)
	writefile := fmt.Sprint(ebt, profit, ratio)

	writeTofile(fileName, writefile)
	value, err := readFromFile(fileName)

	if err != nil {
		fmt.Printf("Errorrr!!!!")
	}

	fmt.Print("value", value)

	fmt.Printf("%.1f\n", ebt)
	fmt.Printf("%.1f\n", profit)
	fmt.Printf("%.3f\n", ratio)
}

func writeTofile(filename string, value string) {
	text := fmt.Sprint(value)
	os.WriteFile(filename, []byte(text), 0644)
}

func readFromFile(filename string) ([]string, error) {
	data, err := os.ReadFile(filename)

	if err != nil {
		return nil, errors.New("file now found")
	}

	text := string(data)

	result := strings.Split(text, " ")

	return result, nil

}

func calculateFinancials(revenue, expenses, taxRate float64) (float64, float64, float64) {
	ebt := revenue - expenses
	profit := ebt * (1 - taxRate/100)
	ratio := ebt / profit
	return ebt, profit, ratio
}

func getUserInput(infoText string) (float64, error) {
	var userInput float64
	fmt.Print(infoText)
	fmt.Scan(&userInput)

	if userInput < 0 {
		return 0, errors.New("value must be greater than 0")
	}

	return userInput, nil
}
