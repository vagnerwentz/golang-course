package main

import "fmt"

func main() {
	// Loop - init, condition, post
	for i := 0; i <= 5; i++ {
		fmt.Printf("%v ", i)
	}

	// Loop - nesting loops
	for i := 0; i <= 5; i++ {
		fmt.Printf("First loop %v \n", i)
		for j := 0; j <= 5; j++ {
			fmt.Printf("Second loop %v \n", j)
		}
	}

	// Loop - break & continue
	x := 0
	for {
		x++
		if x > 8 {
			break
		}
		if x%2 != 0 {
			continue
		}
		fmt.Println(x)
	}
	fmt.Println("Done.")

	// Loop - printing ascii
	for i := 33; i <= 122; i++ {
		fmt.Printf("%v - %#U\n", i, i)
	}
}
