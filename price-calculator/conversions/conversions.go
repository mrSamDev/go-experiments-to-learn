package conversions

import (
	"errors"
	"strconv"
)

func StringsToFloats(stringSlice []string) ([]float64, error) {

	var floatSlice []float64
	for _, str := range stringSlice {
		floatValue, err := strconv.ParseFloat(str, 64)
		if err != nil {
			return nil, errors.New("conversion error: " + err.Error())
		}
		floatSlice = append(floatSlice, floatValue)
	}
	return floatSlice, nil
}
