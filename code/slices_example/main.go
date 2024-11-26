package main

import "fmt"

// slices can have items that are all of the same data type but the size of the slice is not fixed
// Arrays have fixed length of items of the same data type
func main() {
	cards := newDeck()
	fmt.Println(cards.toString())

}

func newCard() string {
	return "three"
}
