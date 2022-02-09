package main

import "fmt"

func main() {
	firstName := make([]string, 0, 3)
	fmt.Printf("The length and the cap of firstName = %v %v\n", len(firstName), cap(firstName))

	lastName := make([]string, 0, 3)
	fmt.Printf("The length and the cap of lastName = %v %v\n", len(lastName), cap(lastName))

	firstName = append(firstName, "Vagner", "Bruna", "Antonio")
	fmt.Printf("The length and the cap of firstName = %v %v\n", len(firstName), cap(firstName))

	lastName = append(lastName, "Wentz", "Eloisa", "Roberval")
	fmt.Printf("The length and the cap of lastName = %v %v\n", len(lastName), cap(lastName))

	fmt.Println(firstName)
	fmt.Println(lastName)

	multiDimensional := [][]string{firstName, lastName}

	fmt.Println(multiDimensional)

	for i, unidimensional := range multiDimensional {
		fmt.Printf("Record: %v\n", i)
		for j, v := range unidimensional {
			fmt.Printf("\t[%v]=%v\n", j, v)
		}
	}
}
