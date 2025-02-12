package fileutils

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func ReadFloatFromFile(fileName string) (float64, error) {
	data, err := os.ReadFile(fileName)

	if err != nil {
		return 1000, errors.New("failed to find  file")
	}

	text := string(data)
	value, err := strconv.ParseFloat(text, 64)

	if err != nil {
		return 1000, errors.New("failed to parse stored value")
	}

	return value, nil
}

func WriteBalanceToFile(fileName string, balance float64) {
	balanceText := fmt.Sprint(balance)
	os.WriteFile(fileName, []byte(balanceText), 0644)
}
