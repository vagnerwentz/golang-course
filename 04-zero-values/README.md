<h1>Zero values</h1>

Basically the zero value is the default value if we do not assign some value to variable.

* false for booleans
* 0 for integers
* 0.0 for floats
* "" for strings
* nil for
  * pointers
  * functions
  * slices
  * channels
  * maps
* use short declaration as much as possible
* use var for
  * zero value
  * package scope
```go
package main

import "fmt"

func main() {
	var boolZeroValue bool
	var intZeroValue int
	var floatZeroValue float64
	var stringZeroValue string

	fmt.Println(boolZeroValue)
	fmt.Println(intZeroValue)
	fmt.Println(floatZeroValue)
	fmt.Println(stringZeroValue)
}

```