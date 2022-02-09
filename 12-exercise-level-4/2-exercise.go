package main

import "fmt"

func main() {
	si := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	for i, v := range si {
		fmt.Printf("[%v]=%v\n", i, v)
	}

	fmt.Printf("%T\n", si)
}
