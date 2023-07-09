package main

import "fmt"

func main() {
	mySlice := []string{"Hi", "There", "You"}

	updateSlice(mySlice)

	fmt.Println(mySlice) // will print [Bye There You] even though we are not passing a reference/address
}

func updateSlice(s []string) {
	s[0] = "Bye"
}
