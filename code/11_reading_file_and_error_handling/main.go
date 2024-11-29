package main

// pass file name
func main() {
	cards := newDeckFromFile("my_cards")
	cards.print()
}

func newCard() string {
	return "three"
}
