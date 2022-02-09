package main

import "fmt"

type vehicle struct {
	doors  int
	colors string
}

type truck struct {
	vehicle
	fourWheel bool
}

type sedan struct {
	vehicle
	luxury bool
}

func main() {

	t := truck{
		vehicle: vehicle{
			doors:  4,
			colors: "black",
		},
		fourWheel: true,
	}

	s := sedan{
		vehicle: vehicle{
			doors:  2,
			colors: "white",
		},
		luxury: true,
	}

	fmt.Println(t)
	fmt.Println(s)
}
