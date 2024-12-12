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

func main() {

	premiumJohhny := newUser("johhny", "premium")
	normalSilke := newUser("Silke", "")

	fmt.Println("User:", premiumJohhny.Name, ", Limit:", premiumJohhny.MessageCharLimit)
	fmt.Println("User:", normalSilke.Name, ", Limit:", normalSilke.MessageCharLimit)
}
