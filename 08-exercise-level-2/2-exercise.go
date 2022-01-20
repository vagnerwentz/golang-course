package main

import "fmt"

func main() {
	a := (42 == 42)
	b := 42 >= 41
	c := 42 <= 43
	d := 42 != 41
	e := 41 > 40
	f := 30 > 35

	fmt.Println(a, b, c, d, e, f)
}
