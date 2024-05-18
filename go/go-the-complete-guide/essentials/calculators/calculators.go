package main

import (
	"fmt"
	"math"
)

func InvestmentCalculator() {
	const inflationRate = 2.5
	var investmentAmount float64
	var years float64
	var expectedReturnRate float64

	fmt.Print("Ivestment Amount: ")
	fmt.Scan(&investmentAmount)

	fmt.Print("Expected Return rate: ")
	fmt.Scan(&expectedReturnRate)

	fmt.Print("years: ")
	fmt.Scan(&years)

	futureValue, futureRealValue := calculateFutureValues(investmentAmount, expectedReturnRate, years, inflationRate)

	fmt.Println(futureValue)
	fmt.Println(futureRealValue)

	// formattedFV := fmt.Sprintf("Future Value: %.1f\n", futureValue)

	// fmt.Print(formattedFV)

	// fmt.Printf(`
	// 	the future value is: %.2f
	// 	the future realValue is: %.2f
	// `, futureValue, futureRealValue)

}

func ProfitCalculator() {
	var revenue float64
	var expenses float64
	var taxRate float64

	fmt.Print("Revenue: ")
	fmt.Scan(&revenue)

	fmt.Print("Expenses: ")
	fmt.Scan(&expenses)

	fmt.Print("Tax Rate: ")
	fmt.Scan(&taxRate)

	ebt := revenue - expenses
	profit := ebt * (1 - taxRate/100)
	ratio := ebt / profit

	fmt.Printf("calculated ebt is %.2f, expenses %.2f and ratio %.2f", ebt, expenses, ratio)
}

func outputText(text, text2 string) {
	fmt.Print(text)
	fmt.Print(text2)
}

func calculateFutureValues(investmentAmount, expectedReturnRate, years, inflationRate float64) (fv float64, rfv float64) {
	fv = investmentAmount * math.Pow((1+expectedReturnRate/100), years)
	rfv = fv / math.Pow(1+inflationRate/100, years)
	//return fv, rfv
	return
}
