package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type bot interface {
	getGreeting() string
}

type englishBot struct {
}

type spanishBot struct {
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

func (s spanishBot) getGreeting() string {
	return "Hola!"
}

func (e englishBot) getGreeting() string {
	return "Hi there!"
}

type logWriter struct{}

func (logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	fmt.Println("Just wrote this many bytes:", len(bs))
	return len(bs), nil
}

func main() {
	fmt.Println("# Using Interfaces")
	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)

	fmt.Println("# Http Client")
	resp, err := http.Get("https://google.com")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	// resp.Body.Read(bs)
	// fmt.Println(string(bs))
	// io.Copy(os.Stdout, resp.Body)
	lw := logWriter{}
	io.Copy(lw, resp.Body)

}
