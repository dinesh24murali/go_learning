package main

// slices can have items that are all of the saem data type
func main() {
	cards := newDeck()
	cards.saveToFile("my_cards")
}

func newCard() string {
	return "three"
}
