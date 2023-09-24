package helpers

import (
	"fmt"
	"strings"
)

// Structure of an animal, exposed to outside packages
type Animal struct {
	Name           string
	Breed          string
	NumberOfLegs   int
	scientificName string //private to the package
}

// Public function that prints the full details of an animal
func (a *Animal) PrintInfo() {
	fmt.Println("Name:", a.Name)
	fmt.Println("Breed:", a.Breed)
	fmt.Println("Number of Logs:", a.NumberOfLegs)
	if strings.EqualFold(a.Breed, "dog") {
		a.scientificName = "Canis"
		fmt.Println("Science calls it a:", a.scientificName)
	}
}
