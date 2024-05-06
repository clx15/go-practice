package main

import "fmt"

//var revenue, expenses, taxRate float64

func main() {

	//fmt.Print("Input your Revenue: ")
	//fmt.Scan(&revenue)

	//fmt.Print("Input your Expenses: ")
	//fmt.Scan(&expenses)

	//fmt.Print("Input your TaxRate: ")
	//fmt.Scan(&taxRate)
	//getRevenue()
	//getExpenses()
	//getTaxRate()

	revenue := getUserInput("Revenue: ")
	expenses := getUserInput("Expenses: ")
	taxRate := getUserInput("Tax Rate: ")

	earningsBeforeTax, earningsAfterTax, ratio := calculations(revenue, expenses, taxRate)

	fmt.Printf("Your earnings are: %.1f\n", earningsBeforeTax)
	fmt.Printf("Your profit is: %.1f\n", earningsAfterTax)
	fmt.Printf("Your profit ratio is: %.1f\n", ratio)

}

func calculations(revenue float64, expenses float64, taxRate float64) (ebt float64, eat float64, ratio float64) {
	ebt = revenue - expenses
	eat = ebt * (taxRate / 100)
	ratio = ebt / eat
	return ebt, eat, ratio
}

//func getRevenue() float64 {
//	fmt.Print("Input your Revenue: ")
//	fmt.Scan(&revenue)
//	return revenue
//}

//func getExpenses() float64 {
//	fmt.Print("Input your Expenses: ")
//	fmt.Scan(&expenses)
//	return expenses
//}
//func getTaxRate() float64 {
//	fmt.Print("Input your TaxRate: ")
//	fmt.Scan(&taxRate)
//	return taxRate
//}

func getUserInput(infoText string) float64 {
	var userInput float64
	fmt.Print(infoText)
	fmt.Scan(&userInput)
	return userInput
}
