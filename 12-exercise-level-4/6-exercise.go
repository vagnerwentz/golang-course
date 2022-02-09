package main

import "fmt"

func main() {
	sliceOfStrings := make([]string, 0, 50)
	fmt.Printf("The cap of sliceOfStrings = %v\n", cap(sliceOfStrings))
	fmt.Printf("The length of sliceOfStrings = %v\n", cap(sliceOfStrings))

	sliceOfStates := []string{
		"Alabama", "Alaska", "Arizona", "Arkansas", "California", "Colorado", "Connecticut", "Delaware", "Florida", "Georgia", "Hawaii", "Idaho", "Illinois", "Indiana", "Iowa", "Kansas", "Kentucky", "Louisiana", "Maine", "Maryland", "Massachusetts", "Michigan", "Minnesota", "Mississippi", "Missouri", "Montana", "Nebraska", "Nevada", "New Hampshire", "New Jersey", "New Mexico", "New York", "North Carolina", "North Dakota", "Ohio", "Oklahoma", "Oregon", "Pennsylvania", "Rhode Island", "South Carolina", "South Dakota", "Tennessee", "Texas", "Utah", "Vermont", "Virginia", "Washington", "West Virginia", "Wisconsin", "Wyoming",
	}

	sliceOfStrings = append(sliceOfStrings, sliceOfStates...)

	fmt.Println(sliceOfStrings)

	fmt.Printf("The cap of sliceOfStrings = %v\n", cap(sliceOfStrings))
	fmt.Printf("The length of sliceOfStrings = %v\n", cap(sliceOfStrings))

	//x := make([]string, 5, 5)
	//x = []string{"One", "Two", "Three", "Four"}
	//
	//fmt.Println(cap(x))
	//fmt.Println(len(x))
	//
	//x = append(x, "five")
	//
	//fmt.Println(cap(x))
	//fmt.Println(len(x))

}
