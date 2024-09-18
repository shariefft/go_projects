package main

import (
	"fmt"
	"struct/helper"
)

func main() {
	noteTitle, error := helper.GetUserInput("Title: ")

	if error != nil {
		fmt.Println("Error:", error)
		return
	}
	

	noteContent, error := helper.GetUserInput("Content: ")

	if error != nil {
		fmt.Println("Error:", error)
		return
	}

	entryOne := helper.NewNote(noteTitle, noteContent)
	entryOne.Display()
	entryOne.Save()

	if error != nil {
		fmt.Println("Saving the note failed")
		return
	}

	fmt.Println("Saving the note succeeded")
}

