<h1>Control flow</h1>

<h2>Loop - init, condition, post</h2>

```
for init; condition; post {

}

for i := 0; i <= 5; i++ {
    fmt.Printf("%v ", i)
}
```

<h2>Loop - nesting loops</h2>
```
for i := 0; i <= 5; i++ {
    fmt.Printf("First loop %v \n", i)
    for j := 0; j <= 5; j++ {
        fmt.Printf("Second loop %v \n", j)
    }
}
```

<h2>Loop - break & continue</h2>

```
x := 0
    for {
        x++
        if x > 8 {
            break
        }
        if x%2 != 0 {
            continue
        }
		fmt.Println(x)
	}
fmt.Println("Done.")
```

<h2>Loop - printing ascii</h2>

```
for i := 33; i <= 122; i++ {
    fmt.Printf("%v - %#U\n", i, i)
}
```

