package main

import (
	"packages-example/helpers"
)

func main() {

	var dog = helpers.Animal{Name: "Shizu", Breed: "Dog", NumberOfLegs: 4}
	dog.PrintInfo()
}
