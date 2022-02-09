package main

import "fmt"

func main() {
	arrayOfInts := [5]int{0, 1, 2, 3, 4}

	for i, v := range arrayOfInts {
		fmt.Printf("[%v]=%v\n", i, v)
	}

	fmt.Printf("%T\n", arrayOfInts)
}
