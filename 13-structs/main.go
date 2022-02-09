package main

import "fmt"

type likes struct {
	like string
}

type person struct {
	firstName string
	lastName  string
	age       int
	likes     likes
}

func main() {

	vagner := person{
		firstName: "Vagner",
		lastName:  "Wentz",
		age:       24,
		likes:     likes{like: "Programação"},
	}

	bruna := person{
		firstName: "Bruna",
		lastName:  "Schvingel",
		age:       21,
		likes:     likes{like: "Nina"},
	}

	fmt.Println(vagner)
	fmt.Println(vagner.firstName, vagner.lastName, vagner.age, vagner.likes)

	fmt.Println(bruna)
	fmt.Println(bruna.firstName, bruna.lastName, bruna.age, bruna.likes)

}
