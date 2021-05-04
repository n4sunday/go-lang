[![Go](https://golang.org/lib/godoc/images/go-logo-blue.svg)](https://golang.org)

### Contents

- [Getstart](#getstart)
- [Intro Go](#intro_go)

---

### <a name="getstart"></a>Getstart

#### Hello Word

```go
package main

import "fmt"

func main() {
	fmt.Println("Hello Go lang")
}
```

#### Go CLI

Complies a bunch of go source code file

```zsh
    go build main.go
```

Complies and executes one or two files

```zsh
    go run main.go
```

Complies and "**install**" a package

```zsh
    go install
```

Downloads the raw source code of someone else's package

```zsh
    go get
```

#### Go Packages

https://golang.org/pkg/

---

<a name="intro_go"></a>

### <a name="intro_go"></a>Intro Go

#### Typed

| Dynamic Types | Static Types |
| :------------ | :----------- |
| Javascript    | C++          |
| Ruby          | Java         |
| Python        | **Go**       |

#### Basic Go Typ

##### Customer tpye declarationes

| Types     | Example             |
| :-------- | :------------------ |
| `bool`    | true false          |
| `string`  | "hello world"       |
| `int`     | 0 -10000 9999       |
| `float64` | 10.0001 0.9 -100.03 |

#### Variable Declarations

```go
func main() {
	var a string = "Ace of Spades"
	fmt.Println(a)

	var b, c int = 1, 2
	fmt.Println(b, c)

	var d = true
	fmt.Println(d)

	var e int
	fmt.Println(e)

	f := "apple"
	fmt.Println(f)
}
```

```zsh
Ace of Spades
1 2
true
0
apple
```

> The **:=** syntax is shorthand for declaring and initializing a variable

#### Basic Functions

```go

func simpleFunc() string {
	return "Five of Diamond"
}
```

#### Basic data structure

##### Array

- Fixed length list of thing

```go
func main() {
	a := []string{"C", "Java", "JavaScript", "Go"}
	fmt.Println(a)
}
```

##### Slices

- An array that can grow or shrink
- Every element in a slice must be of same type

```go
func main() {
	a := []string{"C", "Java", "JavaScript", "Go", simpleFunc()}
	a = append(a, "Python")
	fmt.Println(a)
}

func simpleFunc() string {
	return "Dart"
}
```

##### Loop

Three-component loop

```go
func main() {
	sum := 0
	for i := 1; i < 5; i++ {
		sum += i
	}
	fmt.Println(sum) // 10
}
```

While loop

```go
func main() {
	n := 1
	for n < 5 {
		n *= 2
	}
	fmt.Println(n)
}
```

Infinite loop

```go
func main() {
	sum := 0
	for {
		sum++ // repeated forever
	}
	fmt.Println(sum) // never reached
}
```

For-each range loop

```go
func main() {
	a := []string{"C", "Java", "JavaScript", "Go"}

	for index, item := range a {
		fmt.Println(index, item)
	}
}
```

```zsh
0 C
1 Java
2 JavaScript
3 Go
```

Exit a loop

```go
func main() {
	sum := 0
	for i := 1; i < 5; i++ {
		if i%2 != 0 { // skip odd numbers
			continue
		}
		sum += i
	}
	fmt.Println(sum) // 6 (2+4)
}
```

##### Customer type declaration

`deck.go`

```go
type deck []string

func (d deck) print() {
	for i, item := range d {
		fmt.Println(i, item)
	}
}
```

`main.go`

```go
func main() {
	a := deck{"C", "Java", "JavaScript", "Go"}
	a = append(a, "Dart")
	a.print()
}

```

```zsh
go run main.go deck.go
```

##### Condition
```go
if i%2 == 0 {
	fmt.Printf("%d Even number\n", i)
} else if i%2 != 0 {
	fmt.Printf("%d Odd number\n", i)
}
```

```go
func main() {
	fmt.Println("Hello Go")
	a := "hello"
	if b := check(a); b == "true" {
		fmt.Println("TRUE")
	} else {
		fmt.Println("FALSE")
	}
}

func check(v string) string {
	if v == "hello" {
		return "true"
	} else {
		return "false"
	}
}
```
`Switch Case`
```go
func main() {
	switchNumber(0)
}

func switchNumber(n int) {
	switch n {
	case 0:
		fmt.Println("Zero")
		break
	case 1:
		fmt.Println("One")
		break
	case 2:
		fmt.Println("Two")
		break
	}
}
```


