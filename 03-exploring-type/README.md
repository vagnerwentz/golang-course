<h1>Exploring type</h1>

You can see more about [types](https://golangbyexample.com/all-data-types-in-golang-with-examples/) here.

1. Declare a variable is of a certain TYPE it can only hold VALUES of that type.
2. Primitive data types
   1. In a computer science, primitive data type is either of the following:
      1. **A basic type** is data type provided by a programming language as a basic building block. Most language allow more complicated <i>composite types</i> to be constructed starting from basic types.
      2. **A built-in type** is a data type for which the programming language provides built-in support.
3. Composite data types
   1. In computer science, a composite data type or compound data type is any data type which can be constructed in a program using the programming language's primitive data types and other composite types. It is sometimes called a structure or aggregate data type, although the latter term may also refer to arrays, lists, etc.

[Data types in Go by Geek for Geeks](https://www.geeksforgeeks.org/data-types-in-go/)

```go
package main

import "fmt"

func main() {
	var y = 42
	var message = "Shaken, not stirred"
	var singleChar = 'A'
	var stringLiteral = `James Said, "Shaken, not stirred"`
	var active = true
	var gravity = 9.8

	fmt.Printf("Value of y = %v the type of y is an %T\n", y, y)
	fmt.Printf("Value of message = %v the type of message is %T\n", message, message)
	fmt.Printf("Value of singleChar = %v the type of singleChar is %T\n", singleChar, singleChar)
	fmt.Printf("Value of active = %v the type of active is %T\n", active, active)
	fmt.Printf("Value of gravity = %v the type of gravity is %T\n", gravity, gravity)
	fmt.Printf("Value of stringLiteral = %v the type of stringLiteral is %T\n", stringLiteral, stringLiteral)
}
```