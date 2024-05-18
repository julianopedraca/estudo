package investmentcalculator

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

	futureValue := investmentAmount * math.Pow((1+expectedReturnRate/100), years)
	futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)

	fmt.Println(futureValue)
	fmt.Println(futureRealValue)

	formattedFV := fmt.Sprintf("Future Value: %.1f\n", futureValue)

	fmt.Print(formattedFV)

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
