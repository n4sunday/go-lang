[![Go](https://golang.org/lib/godoc/images/go-logo-blue.svg)](https://golang.org)

### Contents

- [Getstart](#getstart)
- [Intro Go](#intro_go)


<a name="getstart"></a>### Getstart

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

<a name="intro_go"></a>### Intro Go

#### Typed

| Dynamic Types | Static Types |
| ------------- | ------------ |
| Javascript    | C++          |
| Ruby          | Java         |
| Python        | **Go**       |

#### Basic Go Types

| Types   | Example             |
| ------- | ------------------- |
| bool    | true false          |
| string  | "hello world"       |
| int     | 0 -10000 9999       |
| float64 | 10.0001 0.9 -100.03 |

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
The **:=** syntax is shorthand for declaring and initializing a variable
