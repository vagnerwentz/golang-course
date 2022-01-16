package main

import "fmt"

type myOwnType int

var x myOwnType
var yConversion int

func main() {
	fmt.Printf("Variable x has %v value and the type is %T\n", x, x)
	x = 42
	fmt.Printf("Now the value of x is %v\n", x)

	yConversion := int(x)

	fmt.Println(yConversion)
}
