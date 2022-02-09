package main

import "fmt"

func main() {
	subjects := map[string][]string{
		"wentz_vagner": { // []string is redudant use this here
			"Go", "Terere", "Bruna Eloisa",
		},
		"eloisa_bruna": {
			"Moda", "Chá", "Vagner Wentz",
		},
		"nina_wentz_schvingel": {
			"Passear", "Biofresh", "Papai e Mamãe",
		},
	}

	subjects["tesla"] = []string{
		"Mimo", "Dormir", "Brincar",
	}

	for i, v := range subjects {
		fmt.Printf("About the %v\n", i)
		for j, topicAboutSubject := range v {
			fmt.Printf("\t%v - %v\n", j, topicAboutSubject)
		}
	}

	fmt.Println()
	fmt.Println()

	delete(subjects, "wentz_vagner")

	for i, v := range subjects {
		fmt.Printf("About the %v\n", i)
		for j, topicAboutSubject := range v {
			fmt.Printf("\t%v - %v\n", j, topicAboutSubject)
		}
	}
}
