<h1>Short declaration operator and var keyword</h1>

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

<h2>Var keyword</h2>
Other way to create a variable is to use `var` keyword, the `var` way is the same about short declaration operator but just we need to right more code.
The `var` characteristic is equal when we try to change the type of the variable, we can not do this.

```
package main

import "fmt"

func main() {
	var quantityOfStocks = 1000
	var quantityOfKeys int = 35

	fmt.Println(quantityOfStocks)

	quantityOfStocks = 2000

	fmt.Println(quantityOfStocks)
	fmt.Println(quantityOfKeys)
	
	//'"200"' (type string) cannot be represented by the type int
	quantityOfKeys = "200"
}
```

If we want to create a global variable we do not use the `:=` we need to use `var`.
```go
package main

import "fmt"

// # command-line-arguments
//./main.go:4:1: syntax error: non-declaration statement outside function body
quantity := 10
var total int

var amount = 10000

func main() {
	fmt.Println(amount)
	fmt.Println(total)
	foo()
}

func foo() {
	fmt.Println("Amount", amount)
}
```
