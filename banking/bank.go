package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

const accBalFile = "balance.txt"

func saveBalToMemory(balance float64) {
	balanceText := fmt.Sprint(balance)
	os.WriteFile(accBalFile, []byte(balanceText), 0644)
}
func getBalFromMemory() (float64, error) {
	data, err := os.ReadFile(accBalFile)
	if err != nil {
		return 1000, errors.New("Shit's fucked.")
	}
	balanceText := string(data)
	balance, err := strconv.ParseFloat(balanceText, 64)
	if err != nil {
		return 1000, errors.New("Shit's fucked, there's no money")
	}
	return balance, nil
}

func main() {
	var accountBalance, err = getBalFromMemory()

	fmt.Println("Welcome to BankGo")
	if err != nil {
		fmt.Println("Error")
		fmt.Println(err)
		fmt.Println("--------")
		//panic("can't continue lmao")
	}

	for {
		fmt.Println("What would you like to do")
		fmt.Println("1. Check balance")
		fmt.Println("2. Deposit money")
		fmt.Println("3. Withdraw money for Darkutz")
		fmt.Println("4. Exit")

		var choice int
		fmt.Print("What do you want to do: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			fmt.Println("Your balance is: ", accountBalance)
		case 2:
			fmt.Println("Your deposit:")
			var deposit float64
			fmt.Scan(&deposit)
			if deposit <= 0 {
				fmt.Println("Invalid amount. Please deposit an amount greater than 0.")
				continue
			}
			accountBalance += deposit
			fmt.Println("Deposit successful. Your new balance is: ", accountBalance)
			saveBalToMemory(accountBalance)
		case 3:
			fmt.Println("How much money do you want to withdraw for Darkutz's food?")
			var amount float64
			fmt.Scan(&amount)
			if amount <= 0 {
				fmt.Println("Invalid amount. Please withdraw an amount greater than 0.")
				continue
			}
			if amount <= accountBalance {
				accountBalance -= amount
				fmt.Println("Your new balance is: ", accountBalance)
				saveBalToMemory(accountBalance)
			} else {
				fmt.Println("Insufficient funds.")
			}
		default:
			fmt.Println("Exiting..")
			fmt.Println("Thank you for choosing this bank.")
			return
		}
	}

}
