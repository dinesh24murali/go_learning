package main

// slices can have items that are all of the saem data type
func main() {
	cards := newDeck()
	// append will add an item to the slice and return a new slice

	cards.print()
}

func newCard() string {
	return "three"
}
