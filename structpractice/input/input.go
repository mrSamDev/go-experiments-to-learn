package input

import (
	"bufio"
	"errors"
	"fmt"
	"os"
)

// NumberInput is a function.
func Read(prompt string) (string, error) {
	var input string
	fmt.Print(prompt)

	bufio.NewReader(os.Stdin).ReadString('\n')
	if input == "" {
		return "", errors.New("input is required")
	}

	return input, nil
}

// Output is a function.
func Output(text string) {}
