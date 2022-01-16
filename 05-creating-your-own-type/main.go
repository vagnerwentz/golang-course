package main

import "fmt"

var a int

type hotdog int

var b hotdog

func main() {
	a = 42
	b = 43
	fmt.Printf("The value of a is equal to %v and the type of a is an %T\n", a, a)
	fmt.Printf("The value of b is equal to %v and the type of b is a %T\n", b, b)

	/*
		a = b
		b = a
	*/

}
