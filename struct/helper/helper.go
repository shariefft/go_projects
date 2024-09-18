package helper

import (
	"encoding/json"
	"fmt"
	"errors"
	"bufio"
	"os"
	"strings"
)

//define note structs
type Note struct{
	Title string
	Content string
}

//constructor func to create new notes
func NewNote(noteTitle, noteContent string) Note {
	return Note{
		Title: noteTitle,
		Content: noteContent,
	}
  }

//func to get user input
func GetUserInput(textInfo string) (string, error) {
	var userText string
	fmt.Print(textInfo)

	reader := bufio.NewReader(os.Stdin)
	text, error := reader.ReadString('\n')

	if error != nil {
		return "", errors.New("Invalid input")
	}

	userText = strings.TrimSuffix(text, "\n")

	return userText, nil // return the user input
}

// Method to display a note
func (n Note) Display() {
	fmt.Println("Your notes title is: ", n.Title)
	fmt.Println("Your notes content is: ", n.Content)
}


//Method to save a note in a json format
func (note Note) Save() error {
	fileName := strings.ReplaceAll(note.Title, " ", "_")
	fileName = strings.ToLower(fileName) + ".json"

	json, error:= json.Marshal(note)

	if error != nil {
		return error
	}

	return os.WriteFile(fileName, json, 0644)
	
}