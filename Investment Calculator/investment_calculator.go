package main

import (
	"fmt"
	"math"
)

const inflationRate = 2.5

func main() {
	var investmentAmount float64
	var years float64
	var expectedReturnRate float64

	//fmt.Print("Investment Amount: ")
	outputText("Investment Amount: ")
	fmt.Scan(&investmentAmount)

	//fmt.Print("Expected Return Rate: ")
	outputText("Expected Return Rate: ")
	fmt.Scan(&expectedReturnRate)

	//fmt.Print("Years: ")
	outputText("Years: ")
	fmt.Scan(&years)

	futureValue, futureRealValue := calculateFutureValues(investmentAmount, expectedReturnRate, years)
	//futureValue := investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	//futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)

	//formatedFV := fmt.Sprintf("Future Value %.1f\n", futureValue)
	//formatedRFV := fmt.Sprintf("Future Real Value (adjusted for inflation): %.1f\n", futureRealValue)
	// fmt.Println("Future Value:", futureValue)
	//fmt.Printf("Future Value: %v\n", futureRealValue)
	//fmt.Println("Future Real Value (adjusted for inflation):", futureRealValue)
	//fmt.Printf("Future Value: %v\nFuture Real Value (adjusted for inflation): %v\n ", futureRealValue, futureRealValue)
	//fmt.Printf("Future Value: %f\nFuture Real Value (adjusted for inflation): %f\n ", futureRealValue, futureRealValue)
	//fmt.Printf("Future Value: %.1f\nFuture Real Value (adjusted for inflation): %.1f\n ", futureRealValue, futureRealValue)
	//fmt.Print(formatedFV, formatedRFV)
	fmt.Printf(`Future Value: %.1f
	Future Real Value (adjusted for inflation): %.1f\n`, futureValue, futureRealValue)
}

func outputText(text string) {
	fmt.Print(text)
}

func calculateFutureValues(investmentAmount, expectedReturnRate, years float64) (fv float64, rfv float64) {
	fv = investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	rfv = fv / math.Pow(1+inflationRate/100, years)
	//return fv, rfv
	return
}
