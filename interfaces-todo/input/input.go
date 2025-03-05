package input

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func Read(prompt string) string {
	var input string
	fmt.Print(prompt)

	reader := bufio.NewReader(os.Stdin)

	input, err := reader.ReadString('\n')

	if err != nil {
		return ""
	}

	input = strings.TrimSuffix(input, "\n")
	input = strings.TrimSuffix(input, "\r")

	return input
}

// Output is a function.
func Output(text string) {
	fmt.Print(text)
}
