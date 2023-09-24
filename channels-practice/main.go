package main

import (
	"log"
	"math/rand"
	"time"
)

// Print random numbers
// Create function that sets some value to a channel,
// and read the channel back in the main function

// Number pool
const numPool = 1000

func main() {

	intChan := make(chan int)
	defer close(intChan) //Close the channel once the last value for the channel is received

	go getRandomNumber(intChan)
	log.Println(<-intChan)

}

func getRandomNumber(c chan int) {

	randObj := rand.New(rand.NewSource(time.Now().UnixNano())) //Generate random number
	c <- randObj.Intn(numPool)
}
