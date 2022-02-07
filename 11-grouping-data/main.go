package main

import "fmt"

func main() {
	x := []int{1, 2, 3, 4, 5}

	fmt.Printf("The old x = %v\n", x)

	x = append(x, 6)
	fmt.Printf("The x after append x = %v\n", x)

	x = append(x, 7, 8, 9)
	fmt.Printf("The x after append x = %v\n", x)

	y := []int{10, 11, 12, 13}

	z := append(x, y...)

	fmt.Printf("The values of z = %v\n", z)

	// Remove the value 6 and 7
	z = append(z[:5], z[7:]...)

	fmt.Printf("The values of z = %v\n", z)

	a := make([]int, 5, 8)

	fmt.Printf("The values of a=%v\n", a)
	fmt.Printf("The length of a=%v\n", len(a))
	fmt.Printf("The capacity of a=%v\n", cap(a))

	a[0] = 1

	fmt.Printf("The values of a=%v\n", a)

	a = append(a, 10)

	fmt.Printf("The values of a=%v\n", a)

	a = append(a, 11)

	fmt.Printf("The values of a=%v\n", a)

	jb := []string{"James", "Bond", "Chocolate", "Martini"}

	fmt.Println(jb)

	mp := []string{"Miss", "MoneyPenny", "Strawberry", "Hazelnut"}

	fmt.Println(mp)

	xp := [][]string{jb, mp}
	fmt.Println(xp)

	m := map[string]int{
		"Vagner": 1234567,
		"Eloisa": 1234466,
	}

	fmt.Println(m)

	fmt.Println(m["Vagner"])
	fmt.Println(m["Eloisa"])
	fmt.Println(m["Testing"])

	value, ok := m["Testing"]

	fmt.Println(value, ok)

	if !ok {
		fmt.Println("The testing does not exists")
	}

	if v, ok := m["Vagner"]; ok {
		fmt.Println(v)
	}

	m["Todd"] = 9999

	fmt.Println(m)

	for index, value := range m {
		fmt.Println(index, value)
	}

	delete(m, "Vagner")

	fmt.Println(m)

	delete(m, "Baba")

}
