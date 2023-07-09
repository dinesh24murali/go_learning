package main

import "fmt"

func main() {
	// here [string] is for the key, and string is the value
	// colors := map[string]string{
	// 	"red":   "#ff0000",
	// 	"blue":  "#00ff00",
	// 	"green": "#0000ff",
	// }

	// var colors map[string]string

	// this is same as above
	colors := make(map[string]string)

	colors["white"] = "#ffffff"
	colors["red"] = "#ff0000"

	// This doesn't work
	// colors.white

	// Delete an item from the map
	delete(colors, "red")

	fmt.Println(colors)
}
