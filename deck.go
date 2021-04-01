package main

import "fmt"

// Create a new type of 'deck'
// whice is a slice of string
type deck []string

func (d deck) print() {
	for i, item := range d {
		fmt.Println(i, item)
	}
}
