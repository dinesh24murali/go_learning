package main

import (
	"fmt"
	"strings"
)

// Create a new type of 'deck'
// Which is a slice of string

type deck []string

// here "d deck" is referred to as receiver
// This is like a method to the class called deck
func (d deck) print() {
	// naming a variable as underscore (_) will prevent go from throwing error regarding
	// the value is not being used
	for _, card := range d {
		fmt.Println(card)
	}
}

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clover"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

// This is how we get a sub set of the array
// d[0:5] - will return an array that contains the first 5 items
// d[:5] - This is exactly the same as above
// d[5:] - This will return all the items after index 4, not including 4
func deal(d deck, handside int) (deck, deck) {
	return d[:handside], d[handside:]
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}
