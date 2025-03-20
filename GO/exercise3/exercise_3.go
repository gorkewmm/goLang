package main

import (
	"errors"
	"fmt"
	"os"
)

func writeToFile(results1, results2, results3 float64) {
	var str = fmt.Sprintf("EBT:%1.f \nProfit:%1.f\nRatio:%3.f", results1, results2, results3)
	os.WriteFile("calculater.txt", []byte(str), 0644)
}

func main() {
	revenue, err1 := getUserInput("Revenue: ")
	if err1 != nil {
		fmt.Println(err1)
		return
	}
	expenses, err2 := getUserInput("Expenses: ")
	if err2 != nil {
		fmt.Println(err2)
		return
	}
	taxRate, err3 := getUserInput("Tax Rate: ")
	if err3 != nil {
		fmt.Println(err3)
		return
	}

	ebt, profit, ratio := calculateFinancials(revenue, expenses, taxRate)
	writeToFile(ebt, profit, ratio)

	fmt.Printf("%.1f\n", ebt)
	fmt.Printf("%.1f\n", profit)
	fmt.Printf("%.3f\n", ratio)
}

func calculateFinancials(revenue, expenses, taxRate float64) (float64, float64, float64) {
	ebt := revenue - expenses
	profit := ebt * (1 - taxRate/100)
	ratio := ebt / profit
	return ebt, profit, ratio
}

func getUserInput(infoText string) (float64, error) {
	var userInput float64
	fmt.Print(infoText)
	fmt.Scan(&userInput)
	if userInput <= 0 {
		return 0, errors.New("Hatalı giriş")
	}
	return userInput, nil
}
