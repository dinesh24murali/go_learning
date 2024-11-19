package main

import "fmt"

// slices can have items that are all of the saem data type
func main() {
	cards := newDeck()

	fmt.Println(cards.toString())
}
