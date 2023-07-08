package main

import "fmt"

// Go is hardly typed

func main() {
	var card string = "Ace of Spades"
	// This is same as above
	// card1 := "Ace of Spades"
	card = "Ace of Heart"
	fmt.Println(card)
}
