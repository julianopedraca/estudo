package main

import "fmt"

func main() {
	// Arrays
	var productNames [4]string
	productNames = [4]string{"A Book"}
	prices := [4]float64{10.99, 9.99, 45.99, 20.0}
	fmt.Println(prices)

	productNames[2] = "A Carpet"

	fmt.Println(productNames)

	fmt.Println(prices[2])

	//slice
	featuredPrices := prices[1:3]

	fmt.Println(featuredPrices)

	// dynamic lists with slices
	dynamicList := []float64{10.99, 8.99}

	fmt.Println(dynamicList[1])

	updatedDynamicList := append(dynamicList, 5.99)
	fmt.Println(updatedDynamicList)

	unpackingList := []float64{101.99, 80.99, 20.59}
	//unpacking list values
	dynamicList = append(dynamicList, unpackingList...)
	fmt.Println(dynamicList)

}
