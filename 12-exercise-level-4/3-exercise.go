package main

import "fmt"

func main() {
	si := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}

	fmt.Println(si[:5])
	fmt.Println(si[5:])
	fmt.Println(si[2 : len(si)-3])
	fmt.Println(si[1 : len(si)-4])

}
