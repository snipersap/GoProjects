package main

import "fmt"

type messageToSend struct {
	phoneNumber int
	message     string
}

func main() {

	var ms messageToSend
	ms.phoneNumber = 12346
	ms.message = "This is my Phone Number:"

	fmt.Println(ms.message, ms.phoneNumber)

}
