package main

import (
	"fmt"
	"net/http"
	"os"
)

func CheckLink(link string, c chan string) {
	resp, err := http.Get(link)
	if err != nil {
		fmt.Printf("Error for %s; Detail: %s", link, err)
		c <- "Might be down"
		os.Exit(1)
	}
	fmt.Printf("Response status code for %s is %d", link, resp.StatusCode)
	fmt.Println()
	c <- "Yep it is up"
}

func main() {
	links := []string{
		"http://google.com",
		"http://facebook.com",
		"http://stackoverflow.com",
		"http://golang.com",
		"http://amazon.com",
	}

	c := make(chan string)

	for _, link := range links {
		go CheckLink(link, c)
	}
	for _, _ = range links {
		fmt.Println(<-c)
	}
}
