package main

import "fmt"

var wantBeer = false

func main() {
	fmt.Println("Do you want a beer?(true/false)")
	fmt.Scanln(&wantBeer)
	if wantBeer == false {
		fmt.Println("Are you thirsty?(true/false)")
		fmt.Scanln(&wantBeer)
		if wantBeer == false {
			fmt.Println("Is it Saturday?(true/false)")
			fmt.Scanln(&wantBeer)
			if wantBeer == false {
				fmt.Println("Are you sure you don't want a beer?(true/false)")
				fmt.Scanln(&wantBeer)
				getBeer()
			} else {
				getBeer()
			}
		} else {
			getBeer()
		}
	} else {
		getBeer()
	}

}

func getBeer() {
	fmt.Println("Go get a beer")
}
