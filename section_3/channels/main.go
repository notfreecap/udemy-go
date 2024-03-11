package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	links := []string{
		"https://google.com",
		"https://facebook.com",
		"https://stackoverflow.com",
		"https://golang.org",
		"https://amazon.com",
	}

	c := make(chan string)

	for _, link := range links {
		go checkLink(link, c)
	}

	// Using '<-'
	// for i := 0; i < len(links); i++ {
	// 	fmt.Println(<-c)
	// }

	for l := range c {
		go func(link string) {
			time.Sleep(time.Second * 10)
			checkLink(link, c)
		}(l)
	}
}

func checkLink(link string, c chan string) {
	res, err := http.Get(link)

	if err != nil {
		fmt.Println(link, "migth be down!. HTTP STATUS", res.StatusCode)
		c <- link
		return
	}

	fmt.Println(link, "is up!. HTTP STATUS", res.StatusCode)
	c <- link
}
