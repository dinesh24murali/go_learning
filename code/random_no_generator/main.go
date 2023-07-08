package main

// slices can have items that are all of the saem data type
func main() {
	cards := newDeck()
	cards.shuffle()
	cards.print()
}

func newCard() string {
	return "three"
}
