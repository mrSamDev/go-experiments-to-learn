package utils

import (
	"fmt"

	"example.com/app/input"
)

func ReadNote() (string, string, error) {
	title, titleError := input.Read("Note title: ")

	if titleError != nil {
		fmt.Println(titleError)
		return "", "", titleError
	}

	content, contentError := input.Read("Note content: ")

	if contentError != nil {
		fmt.Println(contentError)
		return "", "", contentError
	}

	return title, content, nil
}
