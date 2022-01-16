<h1>Short declaration operator</h1>

You can see more about [short declaration operator](https://www.geeksforgeeks.org/short-variable-declaration-operator-in-go/#:~:text=Short%20Variable%20Declaration%20Operator(%3A%3D)%20in%20Golang%20is%20used,the%20scope%20of%20the%20variables.) here.

Basically this operator could be used to "write" less code.
You create a variable using `:=` and the value you assign this variable will inherit the type of this value and can not be changed in runtime execution.
The value of the variable can be changed to a value with the same type but not different type, to change the value basically you need to use `=` operator after you create a variable using `:=`.

```go
package main

import "fmt"

func main() {
	// The same as var numberOfBytes int = 10
	numberOfBytes := 10

	fmt.Println(numberOfBytes)

	numberOfBytes = 20

	fmt.Println(numberOfBytes)
}
```

```bash
❯ go run main.go
10
20
```

<h2>Using := again with the same variable</h2>
```go
package main

import "fmt"

func main() {
	numberOfBytes := 10

	fmt.Println(numberOfBytes)
	
	//No new variables on the left side of ':='
	numberOfBytes := 30
}
```

<h2>Assigning to other type</h2>
```go
package main

import "fmt"

func main() {
	numberOfBytes := 10

	fmt.Println(numberOfBytes)

	//'"30"' (type string) cannot be represented by the type int
	numberOfBytes = "30"
}
```