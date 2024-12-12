package main

import "fmt"

type Membership struct {
	Type             string
	MessageCharLimit int
}

type User struct {
	Name string
	Membership
}

func newUser(name string, membershipType string) User {

	mcLimit := 100
	if membershipType == "premium" {
		mcLimit = 1000
	}
	newUser := User{
		Name: name,
		Membership: Membership{
			Type:             membershipType,
			MessageCharLimit: mcLimit,
		},
	}
	return newUser

}

func (u User) SendMessage(message string, messageLength int) (string, bool) {
	if messageLength <= u.MessageCharLimit {
		return message, true
	}

	return "", false
}

func main() {

	premiumJohhny := newUser("johhny", "premium")
	normalSilke := newUser("Silke", "")

	fmt.Println("User:", premiumJohhny.Name, ", Limit:", premiumJohhny.MessageCharLimit)
	fmt.Println("User:", normalSilke.Name, ", Limit:", normalSilke.MessageCharLimit)

	messageFromJhonny := "Hello, I am Jhonny"
	_, ok := premiumJohhny.SendMessage(messageFromJhonny, len(messageFromJhonny))
	if ok {
		fmt.Println("Message sent from Jhonny:", messageFromJhonny)
		ok = false
	} else {
		fmt.Println("Message too large to send for ", premiumJohhny.Name)
	}

	messageFromSilke := "Hello dear, I am Silke. How are you today? We are doing fine and very happy! We wanted to wirte to you about our home."
	_, ok = normalSilke.SendMessage(messageFromSilke, len(messageFromSilke))
	if ok {
		fmt.Println("Message sent from Silke:", messageFromSilke)
		ok = false
	} else {
		fmt.Println("Message too large to send for ", normalSilke.Name)
	}

}
