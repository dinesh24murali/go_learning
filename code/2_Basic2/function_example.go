package main

import "fmt"

func main() {
	card := newCard()

	fmt.Println(card)
}

// return type is give after the braces
func newCard() string {
	return "Five of Diamnds"
}
