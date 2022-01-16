package main

import "fmt"

var amount = 10000

func main() {
	numberOfBytes := 10
	fmt.Println(numberOfBytes)

	numberOfBytes = 30
	fmt.Println(numberOfBytes)

	x := numberOfBytes + 10
	fmt.Println(x)

	var quantityOfStocks = 1000
	var quantityOfKeys int = 35

	fmt.Println(quantityOfStocks)

	quantityOfStocks = 2000

	fmt.Println(quantityOfStocks)
	fmt.Println(quantityOfKeys)

	fmt.Println(amount)
	foo()

	//'"200"' (type string) cannot be represented by the type int
	//quantityOfKeys = "200"
}

func foo() {
	fmt.Println("Amount", amount)
}
