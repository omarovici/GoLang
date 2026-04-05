package main

import "fmt"

func main() {
	//var revenue, expenses, taxRate float64
	//fmt.Print("Revenue: ")
	//fmt.Scan(&revenue)
	//fmt.Print("Expenses: ")
	//fmt.Scan(&expenses)
	//fmt.Print("Tax Rate: ")
	//fmt.Scan(&taxRate)
	revenue := textInputUser("Revenue: ")
	expenses := textInputUser("Expenses: ")
	taxRate := textInputUser("Taxes: ")

	//earningsBeforeTax := revenue - expenses
	//profit := earningsBeforeTax * (1 - taxRate/100)
	//ratio := earningsBeforeTax / profit
	earningsBeforeTax, profit, ratio := profitCalculatorValue(revenue, expenses, taxRate)

	fmt.Printf("Earnings before tax: %.1f\n", earningsBeforeTax)
	fmt.Printf("Profit: %.1f\n", profit)
	fmt.Printf("Ratio: %.1f\n", ratio)
}

func profitCalculatorValue(revenue, expenses, taxRate float64) (ebt, profit, ratio float64) {
	ebt = revenue - expenses
	profit = ebt * (1 - taxRate/100)
	ratio = ebt / profit
	return
}

func textInputUser(text string) (value float64) {
	fmt.Print(text)
	fmt.Scan(&value)
	return
}
