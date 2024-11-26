package main

import (
	"fmt"
	"net/http"
)

// This program will keep hitting these URL infinitely
func main() {
	links := []string{"https://google.com", "https://facebook.com", "https://stackoverflow.com", "https://golang.org", "https://amazon.com"}

	c := make(chan string)

	for _, link := range links {
		go checkLink(link, c)
	}
	for {
		// pass the response from the channel which is a string
		go checkLink(<-c, c)
	}
	// This and the above are the same thing
	// for l := range c{
	// 	go checkLink(l, c)
	// }
}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)

	if err != nil {
		fmt.Println(link, "might be down")
		c <- link
		return
	}

	fmt.Println(link, " is Up!")
	c <- link
}
