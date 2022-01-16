<h1>📦 Packages 📦</h1>

To learn more about packages you can read this [article](https://www.callicoder.com/golang-packages/#:~:text=Go%20Package,files%2C%20or%20other%20Go%20packages.&text=The%20above%20package%20declaration%20must,in%20your%20Go%20source%20file.) created by RAJEEV SINGH

In the most basic terms, A package is nothing but a directory inside your Go workspace containing one or more Go source files, or other Go packages.

To see other packages you can see here [Standard Library](https://pkg.go.dev/std).

* Variadic parameters
  * The "...someType" is how we signify a variadic parameter.
  * The type "interface{}" is the empty interface.
    * Everything is of type "interface{}"
  * So the parameter "interface{}" means "give me as many arguments of any type as you would like"
* Throwing away returns
  * use the "_" underscore character to throw away returns
* You can't have unused variable in your code
* We use this notation in Go
  * package.Identifier
    * ex: fmt.Println()

```go
package main

import "fmt"

func main() {
	n, err := fmt.Println("Hello Go", 42, true)
	fmt.Println(n)
	fmt.Println(err)
}
```

```bash
Hello Go 42 true
17
<nil>
```

```go
package main

import "fmt"

func main() {
	n, err := fmt.Println("Hello Go", 42, true)
	fmt.Println(n)
}
```

```bash
❯ go run main.go
# command-line-arguments
./main.go:6:5: err declared but not used
```
