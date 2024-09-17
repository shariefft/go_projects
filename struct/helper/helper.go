package helper

import (
	"fmt"
	"errors"
)

//define note structs
type Note struct{
	title string
	content string
}

//constructor func to create new notes
func NewNote(noteTitle, noteContent string) Note {
	return Note{
		title: noteTitle,
		content: noteContent,
	}
  }

//func to get user input
func GetUserInput(textInfo string) (string, error) {
	var userText string
	fmt.Print(textInfo)
	fmt.Scanln(&userText)

	if userText == "" {
		return "", errors.New("Invalid input")
	}

	return userText, nil // return the user input
}

func structDisplay(n Note) {
	fmt.Println("Your notes title is: ", n.title)
	fmt.Println("Your notes title is: ", n.content)
}