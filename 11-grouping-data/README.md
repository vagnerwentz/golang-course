<h1>Grouping data</h1>

<h2>Array</h2>

You can see more abouts arrays [here](https://go.dev/doc/effective_go#arrays)

Arrays are useful when planning the detailed layout of memory and sometimes can help avoid allocation, but primarily they are a building block for slices, the subject of the next section. To lay the foundation for that topic, here are a few words about arrays.
```go
var x [5]int
fmt.Println(x)
x[3] = 42
fmt.Println(x)
fmt.Println(len(x))

[0 0 0 0 0]
[0 0 0 42 0]
5
```

<h2>Slices</h2>

You can see more abouts slices [here](https://go.dev/doc/effective_go#slices)

Slices wrap arrays to give a more general, powerful, and convenient interface to sequences of data. Except for items with explicit dimension such as transformation matrices, most array programming in Go is done with slices rather than simple arrays.

```go
// Composite Literals - type{values}
x := []int{1, 2, 3, 4, 5} // Slice of ints
fmt.Println(x)

for index, value := range x {
    fmt.Printf("[%v]=%v\n", index, value)
}
```

<h3>Slicing a slice</h3>

```go
x := []int{1, 2, 3, 4, 5}

fmt.Println(x[1:])
fmt.Println(x[1:3])
fmt.Println(x[1 : len(x)-1])

[2 3 4 5]
[2 3]
[2 3 4]

```

<h3>Append to a slice</h3>

```go
x := []int{1, 2, 3, 4, 5}

fmt.Printf("The old x = %v\n", x)

x = append(x, 6)
fmt.Printf("The x after append x = %v\n", x)

x = append(x, 7, 8, 9)
fmt.Printf("The x after append x = %v\n", x)

y := []int{10, 11, 12, 13}

z := append(x, y...)

fmt.Printf("The values of z = %v\n", z)

The old x = [1 2 3 4 5]
The x after append x = [1 2 3 4 5 6]
The x after append x = [1 2 3 4 5 6 7 8 9]
The values of z = [1 2 3 4 5 6 7 8 9 10 11 12 13]
```

<h3>Deleting from a slice</h3>

```go
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

The old x = [1 2 3 4 5]
The x after append x = [1 2 3 4 5 6]
The x after append x = [1 2 3 4 5 6 7 8 9]
The values of z = [1 2 3 4 5 6 7 8 9 10 11 12 13]
The values of z = [1 2 3 4 5 8 9 10 11 12 13]
```

<h3>Make</h3>

You can see more about makes [here](https://go.dev/doc/effective_go#allocation_make)

```go
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

The values of a=[0 0 0 0 0]
The length of a=5
The capacity of a=8
The values of a=[1 0 0 0 0]
The values of a=[1 0 0 0 0 10]
The values of a=[1 0 0 0 0 10 11]
```

<h2>Map</h2>

You can see more about maps [here](https://go.dev/doc/effective_go#maps)

Maps are a convenient and powerful built-in data structure that associate values of one type (the key) with values of another type (the element or value). The key can be of any type for which the equality operator is defined, such as integers, floating point and complex numbers, strings, pointers, interfaces (as long as the dynamic type supports equality), structs and arrays. Slices cannot be used as map keys, because equality is not defined on them. Like slices, maps hold references to an underlying data structure. If you pass a map to a function that changes the contents of the map, the changes will be visible in the caller.

```go

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

map[Eloisa:1234466 Vagner:1234567]
1234567
1234466
0
0 false
The testing does not exists
1234567

```