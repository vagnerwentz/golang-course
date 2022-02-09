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

	mapPerson := map[string]person{
		vagner.lastName: {
			firstName:              vagner.firstName,
			lastName:               vagner.lastName,
			favoriteIceCreamFlavor: vagner.favoriteIceCreamFlavor,
		},
		bruna.lastName: {
			firstName:              bruna.firstName,
			lastName:               bruna.lastName,
			favoriteIceCreamFlavor: bruna.favoriteIceCreamFlavor,
		},
	}

	//mapPerson := map[string]person{
	//	vagner.lastName: vagner
	//	bruna.lastName: bruna
	//}

	for i, v := range mapPerson {
		fmt.Println(i, v)
	}

}
