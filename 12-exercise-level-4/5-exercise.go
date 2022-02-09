package main

import "fmt"

func main() {
	si := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}

	si = append(si[:3], si[6:]...)

	fmt.Println(si)
}
