package main

import (
	"fmt"

	"example.com/price-calculator/fileManager"
	"example.com/price-calculator/prices"
)

func main() {

	taxRate := []float64{0.05, 0.07, 0.09}

	fm := fileManager.NewFileManager("price.txt", "output")
	// cmd := cmdmanager.NewCMDManager()

	for _, taxRate := range taxRate {

		priceJob := prices.NewTaxIncludedPriceJob(fm, taxRate)

		error := priceJob.Process()

		if error != nil {
			fmt.Println("Error processing job:", error)
		}

	}

}
