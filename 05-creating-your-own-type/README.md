<h1>Own type</h1>

You can see more about to create own type [here](https://appdividend.com/2019/03/22/go-custom-type-declarations-tutorial-with-example/).

```go
package main

import "fmt"

var a int

type hotdog int

var b hotdog

func main() {
  a = 42
  b = 43
  fmt.Printf("The value of a is equal to %v and the type of a is an %T\n", a, a)
  fmt.Printf("The value of b is equal to %v and the type of b is a %T\n", b, b)
}
```

<h2>Can not assign the type hotdog to a variable that is an int</h2>

```go
package main

import "fmt"

var a int

type hotdog int

var b hotdog

func main() {
	a = 42
	b = 43
	fmt.Printf("The value of a is equal to %v and the type of a is an %T\n", a, a)
	fmt.Printf("The value of b is equal to %v and the type of b is a %T\n", b, b)
	
	//Cannot use 'b' (type hotdog) as the type in
	//a = b
	
	//Cannot use 'a' (type int) as the type hotdog
	//b = a
}

```