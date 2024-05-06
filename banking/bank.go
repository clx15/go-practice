package main

import "fmt"

func main() {
	var accountBalance float64 = 1000

	fmt.Println("Welcome to BankGo")
	fmt.Println("What would you like to do")
	fmt.Println("1. Check balance")
	fmt.Println("2. Deposit money")
	fmt.Println("3. Withdraw money for Darkutz")
	fmt.Println("4. Exit")

	var choice int
	fmt.Print("What do you want to do: ")
	fmt.Scan(&choice)

	if choice == 1 {
		fmt.Println("Your balance is, ", accountBalance)
	} else if choice == 2 {
		fmt.Println("Your deposit:")
		var deposit float64
		fmt.Scan(&deposit)
		accountBalance += deposit
		fmt.Println("Your new balance is, ", accountBalance)
	} else if choice == 3 {
		fmt.Println("How much money do you want to withdraw for Darkutz's food?")
		var amount float64
		fmt.Scan(&amount)
		if amount <= accountBalance {
			accountBalance -= amount
			fmt.Println("Your new balance is, ", accountBalance)
		} else {
			fmt.Println("Insufficient funds")
		}

	} else {
		fmt.Println("Exiting..")
	}

}
