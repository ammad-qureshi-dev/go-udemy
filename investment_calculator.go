package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Print("investment amount: $")
	var investmentAmount float64
	fmt.Scan(&investmentAmount)

	fmt.Print("expected return rate: ")
	var expectedReturnRate float64
	fmt.Scan(&expectedReturnRate)
	
	fmt.Print("years: $")
	var years float64
	fmt.Scan(&years)
	
	const INFLATION_RATE float64 = 2.5

	var futureValue float64 = investmentAmount * math.Pow(1 + expectedReturnRate / 100, float64(years))
	fvAdjusted := futureValue / math.Pow(1 + INFLATION_RATE / 100, float64(years))
	fmt.Printf("This is the future value of that amount: $%f\n", futureValue)
	fmt.Printf("This is the adjusted value of that amount: $%f\n", fvAdjusted)

 }