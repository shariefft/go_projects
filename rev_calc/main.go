package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	var revenue float64  // Declare revenue as float64
	var expenses float64 // Declare expenses as float64
	var taxRate float64  // Declare taxRate as float64

	revenue, err := getUserInput("Revenue: ") // Assign to revenue and err
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	expenses, err = getUserInput("Expenses: ") // Assign to expenses and err
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	taxRate, err = getUserInput("Tax Rate: ") // Assign to taxRate and err
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	ebt, profit, ratio := calc(revenue, expenses, taxRate)
	writeResult(ebt, profit, ratio)

	fmt.Println(ebt)
	fmt.Println(profit)
	fmt.Println(ratio)
}

func calc(revenue, expenses, taxRate float64) (float64, float64, float64) {
	ebt := revenue - expenses
	profit := ebt * (1 - taxRate/100)
	ratio := ebt / profit
	return ebt, profit, ratio
}

func getUserInput(textInfo string) (float64, error) {
	var userText float64
	fmt.Print(textInfo)
	fmt.Scan(&userText)

	if userText <= 0 {
		return 0, errors.New("Invalid input: Value cannot be negative or zero.")
	}

	return userText, nil // return the user input
}

func formatResult(ebt, profit, ratio float64) string {
	return fmt.Sprintf("ebt: %.2f\nprofit: %.2f\nratio: %.2f", ebt, profit, ratio)
}

func writeResult(ebt, profit, ratio float64) {
	fileContent := formatResult(ebt, profit, ratio)
	os.WriteFile("result.txt", []byte(fileContent), 0644)
	fmt.Print("Successfully wrote result to file")
}
