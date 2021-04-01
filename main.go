package main

import "fmt"

func main() {
	a := []string{"C", "Java", "JavaScript", "Go"}

	for i, item := range a {
		fmt.Println(i, item)
	}
}
