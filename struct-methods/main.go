package main

import "fmt"

type authenticationInfo struct {
	username string
	password string
}

func (a authenticationInfo) getBasicAuth() string {
	output := fmt.Sprintf("Authorization: Basic %s:%s", a.username, a.password)
	return output
}

func main() {
	a := authenticationInfo{
		username: "HarryPotter",
		password: "lumus",
	}

	fmt.Println(a.getBasicAuth())
}
