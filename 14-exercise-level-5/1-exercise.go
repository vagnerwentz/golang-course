package main

import "fmt"

type person struct {
	firstName              string
	lastName               string
	favoriteIceCreamFlavor []string
}

func main() {
	vagner := person{
		firstName:              "Vagner",
		lastName:               "Wentz",
		favoriteIceCreamFlavor: []string{"Chocolate"},
	}

	bruna := person{
		firstName:              "Bruna",
		lastName:               "Eloisa",
		favoriteIceCreamFlavor: []string{"Mix", "Choco"},
	}

	fmt.Printf("The favorite ice cream flavor of %v %v", vagner.firstName, vagner.lastName+" is ")
	for _, v := range vagner.favoriteIceCreamFlavor {
		fmt.Printf("%v\n", v)
	}

	fmt.Printf("The favorite ice cream flavor of %v %v", bruna.firstName, bruna.lastName+" is ")
	for _, v := range bruna.favoriteIceCreamFlavor {
		fmt.Printf("%v\n", v)
	}

}
