package main

import "fmt"

// Create a new type of 'deck'
// Which is a slice of string

type deck []string

// here "d deck" is referred to as receiver
// This is like a method to the class caled deck
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

// Note that the function decleration need not be in this file even though the
// type deck is declared here
