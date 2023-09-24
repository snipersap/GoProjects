package main

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"
)

type Person struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Country   string `json:"country"`
	Age       int    `json:"age"`
}

func main() {
	var p []Person
	iJ := initJson()

	//Decode json into Person
	err := json.Unmarshal([]byte(iJ), &p)
	if err != nil {
		log.Println("JSON Unmarshal error:", err)
	} else {
		fmt.Println("Persons recovered are:", p)
	}

	//Encode Person into json
	jByteSlice, err := json.MarshalIndent(p, "", "	")
	// jByteSlice, err := json.Marshal(p)
	if err != nil {
		log.Println("Person Marshal error:", err)
	} else {
		fmt.Println("Person Marshaled into:\n", string(jByteSlice))
	}

	//Compare the initial and encoded json strings
	if strings.EqualFold(iJ, string(jByteSlice)) {
		fmt.Println("Initial josn and encoded json is the same.")
	} else {
		log.Fatalln("Encoded json does not match initial json!")
	}
}

func initJson() string {
	return `
	[
		{
			"first_name": "Raju",
			"last_name": "Gentleman",
			"country": "India",
			"age": 10
		},
		{
			"first_name": "Miroslav",
			"last_name": "Guzmann",
			"country": "Ukraine",
			"age": 25
		}
	]
	`

}
