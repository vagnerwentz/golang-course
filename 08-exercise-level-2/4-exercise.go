package main

import (
	"fmt"
	"math"
)

/** Assigns an int to a variable
* Prints that int in decimal, binary, and hex
* Shifts the bits of that int over 1 position to the left, and assigns that to a variable
* Prints that variable in decimal, binary, and hex.*/
func main() {
	var number = 42
	fmt.Printf("Decimal %v, Binary %b, Hexadecimal %#x\n", number, number, number)

	fmt.Println("Shift the bits")

	var otherNumber = number << 1

	fmt.Printf("Decimal %v, Binary %b, Hexadecimal %#x\n", otherNumber, otherNumber, otherNumber)

	fmt.Println("Shift the bits")

	var someNumber = number >> 1

	fmt.Printf("Decimal %v, Binary %b, Hexadecimal %#x\n", someNumber, someNumber, someNumber)

	fmt.Println(math.Pow(2, 5))
}
