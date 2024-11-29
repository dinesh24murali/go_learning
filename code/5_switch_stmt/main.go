package main

import (
	"fmt"
)

func main() {
	input := 10

	switch input {
	case 10:
		fmt.Println(11)
	case 11:
		fmt.Println(12)
	}
	testing()
}

func testing() {
	fmt.Print("Go runs on ")
	value11 := 11
	value12 := 12
	switch value12 = value12 + 1; value12 {
	case value11:
		fmt.Println("OS X.")
	case value12:
		fmt.Println("Linux.")
	default:
		fmt.Printf("%d.", value12)
	}
}
