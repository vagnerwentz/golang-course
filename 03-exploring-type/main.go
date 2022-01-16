package main

import "fmt"

func main() {
	var y = 42
	var message = "Shaken, not stirred"
	var singleChar = 'A'
	var stringLiteral = `James Said, "Shaken, not stirred"`
	var active = true
	var gravity = 9.8

	fmt.Printf("Value of y = %v the type of y is an %T\n", y, y)
	fmt.Printf("Value of message = %v the type of message is %T\n", message, message)
	fmt.Printf("Value of singleChar = %v the type of singleChar is %T\n", singleChar, singleChar)
	fmt.Printf("Value of active = %v the type of active is %T\n", active, active)
	fmt.Printf("Value of gravity = %v the type of gravity is %T\n", gravity, gravity)
	fmt.Printf("Value of stringLiteral = %v the type of stringLiteral is %T\n", stringLiteral, stringLiteral)
}
