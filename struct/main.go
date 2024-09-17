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
	fmt.Println('\n')
	fmt.Println(entryOne)

}

