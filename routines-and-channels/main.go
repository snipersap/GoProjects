package main

import (
	"fmt"
	"net/http"
)

func main() {
	links := []string{
		"http://google.com",
		"http://amazon.com",
		"http://honeywell.com",
		"http://arunisgood.com",
	}
	c := make(chan string)

	for _, link := range links {
		go checkLink(link, c)
	}
	fmt.Println(<-c)
}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		println("Maybe", link, "is down!")
		c <- "vvv DOWN vvv"
	} else {
		println("Yay!!", link, "is up.")
		c <- "^^^ UP ^^^"
	}
}
