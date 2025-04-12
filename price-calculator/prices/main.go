package prices

import (
	"fmt"

	iomanager "example.com/price-calculator/ioManager"
)

type TaxIncludedPriceJob struct {
	IOManager         iomanager.IOManager `json:"-"`
	TaxRate           float64             `json:"tax_rate"`
	InputPrices       []float64           `json:"input_prices"`
	TaxIncludedPrices map[string]string   `json:"tax_included_prices"`
}

//to exclude `json:"-"` from the struct

func (job *TaxIncludedPriceJob) GetData() error {

	prices, err := job.IOManager.ReadAndGetPrice()

	if err != nil {
		return err
	}

	job.InputPrices = prices

	return nil

}

func (job *TaxIncludedPriceJob) Process() error {
	error := job.GetData()

	if error != nil {
		fmt.Println("Error reading file:", error)
		return error
	}

	job.TaxIncludedPrices = make(map[string]string)

	for _, price := range job.InputPrices {
		priceWithTax := price + (price * job.TaxRate)
		key := fmt.Sprintf("%.2f", price)
		job.TaxIncludedPrices[key] = fmt.Sprintf("%.2f", priceWithTax)
	}

	outputfilename := job.IOManager.GenerateOutputFilePath(job.TaxRate)

	return job.IOManager.WriteFileWithPrices(outputfilename, job)

}

func NewTaxIncludedPriceJob(fm iomanager.IOManager, taxRate float64) *TaxIncludedPriceJob {
	return &TaxIncludedPriceJob{
		IOManager: fm,
		TaxRate:   taxRate,
	}
}

/*

map[string]float64 = in js {
	"10.50": 0.3
}

map[string][]float64 = in js {
	"10.50": [0.3,4.3,5.3],
}
*/
