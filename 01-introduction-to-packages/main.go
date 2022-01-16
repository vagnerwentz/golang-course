package main

import "fmt"

func main() {
	n, err := fmt.Println("Hello Go", 42, true)
	fmt.Println(n)
	fmt.Println(err)

	nB, _ := fmt.Println("Hi")
	fmt.Println(nB)
}
