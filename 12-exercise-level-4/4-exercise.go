package main

import "fmt"

func main() {
	si := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	si = append(si, 52)
	fmt.Println(si)

	si = append(si, 53, 54, 55)

	fmt.Println(si)

	osi := []int{56, 57, 58, 59, 60}

	si = append(si, osi...)

	fmt.Println(si)
}
