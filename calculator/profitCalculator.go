package main

import (
	"fmt"
)

func main() {
	var revenue float64
	fmt.Print("Revenue: $")
	fmt.Scan(&revenue)

	var expenses float64
	fmt.Print("Expenses: $")
	fmt.Scan(&expenses)

	var taxRate float64
	fmt.Print("taxRate: ")
	fmt.Scan(&taxRate)

	earningsBeforeTax := revenue - expenses
	profit := earningsBeforeTax * (1 - (taxRate / 100))
	ratio := earningsBeforeTax / profit

	fmt.Printf("Earnings Before Tax: $%f\nProfit: $%f\nRatio: %f\n", earningsBeforeTax, profit, ratio)
}