package iomanager

type IOManager interface {
	ReadAndGetPrice() ([]float64, error)
	GenerateOutputFilePath(taxRate float64) string
	WriteFileWithPrices(filename string, data any) error
}
