package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.org",
		"http://amazon.com",
		"http://defnitelyfakeurl.com",
	}

	// Create a channel to communicate between goroutines
	c := make(chan string)

	// Initial loop to start a goroutine for each link
	for _, link := range links {
		go checkLink(link, c)
	}

	// Infinite loop waiting for a channel message to return
	for l := range c {
		// Spawn an anonymous function (function literal) to handle the delay
		go func(link string) {
			time.Sleep(5 * time.Second)
			checkLink(link, c)
		}(l) // Pass 'l' by value to avoid referencing the same memory address in the loop
	}
}

func checkLink(link string, c chan string) {
	// Make a blocking HTTP GET request inside this specific goroutine
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down!")
		c <- link // Send the link back into the channel
		return
	}

	fmt.Println(link, "is up!")
	c <- link // Send the link back into the channel
}
