package cmdmanager

import (
	"fmt"

	"example.com/price-calculator/conversions"
)

type CMDManager struct {
}

func (fm CMDManager) GenerateOutputFilePath(id float64) string {
	return fmt.Sprintf("output_%v.json", id)
}

func (cmd CMDManager) ReadAndGetPrice() ([]float64, error) {
	fmt.Println("Enter prices")
	var inputs []string
	for {
		fmt.Print("Enter price (or 'done' to finish): ")

		var price string
		fmt.Scanln(&price)

		if price == "done" {
			break
		}

		inputs = append(inputs, price)

	}

	return conversions.StringsToFloats(inputs)

}

func (cmd CMDManager) WriteFileWithPrices(outputFilePath string, data any) error {

	fmt.Println("Writing to file:", outputFilePath)
	fmt.Println("file:", data)

	return nil

}

func NewCMDManager() CMDManager {
	return CMDManager{}
}
