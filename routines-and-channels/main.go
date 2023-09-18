package main

import (
	"net/http"
	"time"
)

func main() {
	links := []string{
		"http://google.com",
		"http://amazon.com",
		"http://honeywell.com",
		"http://arunisgood.com",
	}
	// Create channel
	c := make(chan string)

	// Call routine and pass each link to it
	for _, link := range links {
		go checkLink(link, c)
	}

	//Keep checking the status of the links endlessly
	// for { //endless loop
	// 	go checkLink(<-c, c) //use the value of channel c to check the link again
	// }

	//Alternative form of endless loop checking link status endlessly
	//Include sleep and remove sharing of variables
	for l := range c { //still and endless loop
		// time.Sleep(5 * time.Second)  //Sleeps the main routine
		go func(link string) { //Function literal with func definition
			time.Sleep(5 * time.Second)
			checkLink(link, c)
		}(l) //With actual arguments
	}

	// Wait for the completion of the routines of each of the links
	// for i := 0; i < len(links); i++ {
	// 	fmt.Println(<-c)
	// }

	// Wait using separate Println statements
	// fmt.Println(<-c)
	// fmt.Println(<-c)
	// fmt.Println(<-c)
	// fmt.Println(<-c)
	// fmt.Println(<-c)
}

func checkLink(link string, c chan string) {
	// time.Sleep(5 * time.Second) //sleeps the child routine, but also breaks the single responsiblity principle
	_, err := http.Get(link)
	if err != nil {
		println("Maybe", link, "is down!")
		c <- link
	} else {
		println("Yay!!", link, "is up.")
		c <- link
	}
}
