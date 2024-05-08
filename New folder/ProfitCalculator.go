package main

import (
	"fmt"
	"os"
)

const businessData = "Businesslog.txt"

func saveLogToMemory(sEbt float64, sEat float64, sRatio float64) {
	logText := fmt.Sprintf("Your earnings are: %.1f\nYour profit is: %.1f\nYour profit ratio is: %.1f\n", sEbt, sEat, sRatio)
	//logText = fmt.Sprint("Your profit is: %.1f\n", seat)
	//logText = fmt.Sprint("Your profit ratio is: %.1f\n", sratio)
	os.WriteFile(businessData, []byte(logText), 0644)
}

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

	if revenue > 0 && expenses > 0 && taxRate > 0 {
		earningsBeforeTax, earningsAfterTax, ratio := calculations(revenue, expenses, taxRate)

		fmt.Printf("Your earnings are: %.1f\n", earningsBeforeTax)
		fmt.Printf("Your profit is: %.1f\n", earningsAfterTax)
		fmt.Printf("Your profit ratio is: %.1f\n", ratio)
	} else {
		fmt.Println("The values provided must be greater than zero. Please try again")

	}

}

func calculations(revenue float64, expenses float64, taxRate float64) (ebt float64, eat float64, ratio float64) {
	ebt = revenue - expenses
	eat = ebt * (taxRate / 100)
	ratio = ebt / eat
	saveLogToMemory(ebt, eat, ratio)
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
