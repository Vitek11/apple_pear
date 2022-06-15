package main

import "fmt"

func main() {

	var (
		applePrice = 5.99
		pearPrice  = 7.
		money      = 23.
		numApple   = 9
		numPear    = 8
		totalCost  float64
	)

	fmt.Println("Apple price:", applePrice, "grn")
	fmt.Println("Pear price:", pearPrice, "grn")
	fmt.Println("We have:", money, "grn")

	totalCost = float64(numApple)*applePrice + float64(numPear)*pearPrice
	fmt.Println("1. To buy", numApple, "apples and", numPear, "pears you need", totalCost)
	fmt.Println("2. We can buy", int32(money/applePrice), "apples")
	fmt.Println("3. We can buy", int32(money/pearPrice), "pears")
	numApple = 2
	numPear = 2
	totalCost = float64(numApple)*applePrice + float64(numPear)*pearPrice
	fmt.Println("4. Can we buy", numApple, "apples and", numPear, "pears?", totalCost < money)
}
