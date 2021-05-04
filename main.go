package main

import "fmt"

func main() {
	fmt.Println("Hello Go")
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
