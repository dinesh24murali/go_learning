package main

import (
	"fmt"
	"net/http"
	"time"
)

// This program will keep hitting these URL infinitely
func main() {
	links := []string{"https://google.com", "https://facebook.com", "https://stackoverflow.com", "https://golang.org", "https://amazon.com"}

	c := make(chan string)

	for _, link := range links {
		go checkLink(link, c)
	}
	for l := range c {

		fmt.Println(" IS this printing " + l)

		go func(n string) {
			time.Sleep(time.Second * 5)
			checkLink(n, c)
		}(l)
	}
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
