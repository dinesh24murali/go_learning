package main

import "fmt"

// Slices don't have fixed length
// Internally slice uses array under the hood. It is like a struct
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
