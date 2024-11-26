package main

import (
	"fmt"
	"net/http"
)

// Channels have type string, int, float
func main() {
	links := []string{"https://google.com", "https://facebook.com", "https://stackoverflow.com", "https://golang.org", "https://amazon.com"}

	c := make(chan string)

	for _, link := range links {
		go checkLink(link, c)
	}
	// waiting for all the child routines to end
	for i := 0; i < len(links); i++ {
		fmt.Println(<-c)
	}
}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)

	if err != nil {
		fmt.Println(link, "might be down")
		c <- "Might be down"
		return
	}

	fmt.Println(link, " is Up!")
	c <- "It is up"
}
