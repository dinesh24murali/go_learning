package main

import "fmt"

// Arrays have fixed length. They need to have the same data type
func main() {
	c := [2]string{"one", "two"}
	// This is equal to above
	// c := [...]string{"one", "two"}

	fmt.Printf("%T", c)

	for _, item := range c {
		fmt.Println(item)
	}
}
