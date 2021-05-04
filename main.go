package main

import "fmt"

func main() {
	msg := "some message"
	var msgPointer *string = &msg
	fmt.Println(msg)
	fmt.Println(msgPointer)
	// return 0xc000050240
	fmt.Println(*msgPointer)
	// return "some message"
}
