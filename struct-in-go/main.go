package main

import "fmt"

type messageToSend struct {
	phoneNumber type int
	message type string
}

func main {

	var ms type messageToSend
	ms.phoneNumber = 12346
	ms.message = "This is my Phone Number:"

	fmt.Println(ms.phoneNumber, ms.message)

}