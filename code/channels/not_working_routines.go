package main

import (
	"fmt"
	"net/http"
)

// With this implementation, the main routine will end before the child routines
// end, so we won't see any output. Channels help to solve this
func main() {
	links := []string{"https://google.com", "https://facebook.com", "https://stackoverflow.com", "https://golang.org", "https://amazon.com"}

	for _, link := range links {
		// will span a child routine
		go checkLink(link)
	}
}

func checkLink(link string) {
	_, err := http.Get(link)

	if err != nil {
		fmt.Println(link, "might be down")
		return
	}

	fmt.Println(link, " is Up!")
}
