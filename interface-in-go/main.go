package main

import "fmt"

type bot interface {
	greeting() string
}
type englishBot struct{}
type deutscheBot struct{}

func main() {

	eb := englishBot{}
	db := deutscheBot{}

	print(eb)
	print(db)
}

func print(b bot) {
	fmt.Println(b.greeting())
}

func (englishBot) greeting() string {
	return "Hello there!"
}

func (deutscheBot) greeting() string {
	return "Halli Hallo!"
}

// Print englishBot greeting
// func print(eb englishBot) {
// 	fmt.Println(eb.greeting())
// }

// Print deutscheBot greeting
// func print(db deutscheBot) {
// 	fmt.Println(db.greeting())
// }
