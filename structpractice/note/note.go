package note

import (
	"errors"
	"fmt"
	"time"
)

type Note struct {
	title     string
	content   string
	createdAt time.Time
}

func (note Note) Display() {
	fmt.Printf("Your entered note title %v with content \n\n%v", note.title, note.content)
	// fmt.Println(note)
}

func New(title, content string) (Note, error) {

	if title == "" {
		return Note{}, errors.New("title is required")
	}

	if content == "" {
		return Note{}, errors.New("content is required")
	}

	return Note{
		title:     title,
		content:   content,
		createdAt: time.Now(),
	}, nil
}
