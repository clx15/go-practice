package main

import (
	"fmt"
	"math"
)

const inflationRate = 2.5

func main() {
	var investAmount, years, expectReturnRate float64

	fmt.Print("enter the invest amount: ")
	fmt.Scan(&investAmount)

	fmt.Print("enter the years: ")
	fmt.Scan(&years)

	fmt.Print("enter the expected return rate: ")
	fmt.Scan(&expectReturnRate)

	futureValue, futureRealValue := futureValues(investAmount, expectReturnRate, years)

	formattedFV := fmt.Sprintf("The future value is: %.1f\n", futureValue)
	formattedFRV := fmt.Sprintf("The future real value is: %.1f\n", futureRealValue)

	//	fmt.Printf(`The future value is: %.0f
	//Adjusted for inflation: %.0f`, futureValue, futureRealValue)
	fmt.Print(formattedFV, formattedFRV)

}

func futureValues(investAmount float64, expectReturnRate float64, years float64) (fv float64, rfv float64) {
	fv = investAmount * math.Pow(1+expectReturnRate/100, years)
	rfv = fv / math.Pow(1+inflationRate/100, years)
	return fv, rfv
}
