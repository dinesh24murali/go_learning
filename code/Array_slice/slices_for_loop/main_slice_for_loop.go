package main

import "fmt"

// slices can have items that are all of the saem data type
func main() {
	cards := []string{"one", newCard()}
	// append will add an item to the slice and return a new slice
	cards = append(cards, "four")

	for i, card := range cards {
		fmt.Println(i, card)
	}
}

func newCard() string {
	return "three"
}
