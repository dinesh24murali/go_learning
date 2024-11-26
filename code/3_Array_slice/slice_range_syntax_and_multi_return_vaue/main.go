package main

import "fmt"

// slices can have items that are all of the saem data type
func main() {
	cards := newDeck()
	// append will add an item to the slice and return a new slice
	hand, remainingDeck := deal(cards, 5)
	hand.print()
	fmt.Println("========================================================")
	remainingDeck.print()
}

func newCard() string {
	return "three"
}
