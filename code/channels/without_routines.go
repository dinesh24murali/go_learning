package main

import (
	"fmt"
	"net/http"
)

func main() {
	links := []string{"https://google.com", "https://facebook.com", "https://stackoverflow.com", "https://golang.org", "https://amazon.com"}

	for _, link := range links {
		checkLink(link)
	}
}

func checkLink(link string) {
	_, err := http.Get(link) // <---- this will be a blocking call

	if err != nil {
		fmt.Println(link, "might be down")
		return
	}

	fmt.Println(link, " is Up!")
}
