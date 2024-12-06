package main

import "fmt"

type messageToSend struct {
	message   string
	sender    user
	recipient user
}

type user struct {
	name   string
	number int
}

func canSendMessage(mToSend messageToSend) bool {
	//validate sender
	if mToSend.sender.name == "" || mToSend.sender.number == 0 {
		return false
	}
	//validate receiver
	if mToSend.recipient.name == "" || mToSend.recipient.number == 0 {
		return false
	}

	return true
}

func main() {
	msg := messageToSend{message: "hi"}
	result := canSendMessage(msg)
	fmt.Println("Result:", result)
}
