package utils

import (
	"example.com/app/input"
)

func ReadNote() (string, string) {
	title := input.Read("Note title: ")

	content := input.Read("Note content: ")

	return title, content
}
