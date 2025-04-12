package fileManager

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strconv"
)

type FileManager struct {
	InputFilePath    string
	OutputFilePrefix string
}

func (fm FileManager) GenerateOutputFilePath(id float64) string {
	return fmt.Sprintf("%v_%v.json", fm.OutputFilePrefix, id)
}

func NewFileManager(inputFilePath, outputFilePrefix string) FileManager {
	return FileManager{
		InputFilePath:    inputFilePath,
		OutputFilePrefix: outputFilePrefix,
	}
}

func (fm FileManager) ReadAndGetPrice() ([]float64, error) {
	file, error := os.Open(fm.InputFilePath)

	if error != nil {

		return nil, errors.New("error in opening file")
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	var priceList []float64

	for scanner.Scan() {
		line := scanner.Text()
		var price float64

		price, err := strconv.ParseFloat(line, 64)

		if err != nil {
			fmt.Println("Error reading price:", err)
			continue
		}
		priceList = append(priceList, price)

	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return nil, errors.New("scanner failed to scan file")
	}

	return priceList, nil

}

func (fm FileManager) WriteFileWithPrices(outputFilePath string, data any) error {

	file, error := os.Create(outputFilePath)

	if error != nil {
		return errors.New("file creation error")
	}

	defer file.Close()

	jsonEncoder := json.NewEncoder(file)

	error = jsonEncoder.Encode(data)

	if error != nil {
		return errors.New("data filed converted to json")
	}

	return nil
}
