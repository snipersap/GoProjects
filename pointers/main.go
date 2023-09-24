package main

import (
	"fmt"
	"log"
	"sort"
	"time"
)

// global type available only inside main package
type deepColor struct {
	color string
}

// Struct available outside of package (if package was not main)
type User struct {
	FirstName, LastName, PhoneNumber string //Fields available outside of package
	personnelID                      int    //Private field not available outside of package even if User is exposed to other packages
}

func main() {

	var color string
	color = "Red"
	fmt.Println("Pointers with standard type string >>")
	log.Println("Current color:", color)
	myFavColor(&color) //need to pass the address of the var
	log.Println("Fav color:", color)

	fmt.Println("Pointers with struct>>")
	dColor := deepColor{"Orange"}
	log.Println("Current color:", dColor.color)
	myfavDeepColor(&dColor)
	log.Println("Fav color:", dColor.color)

	log.Println(time.Millisecond)

	// Slices and sorting
	names := []string{"John", "Amar", "Yogi"}
	ages := []int{1, 4, 6, 23, 54, 12, 22}
	names = append(names, "Iolo!")
	ages = append(ages, 0)

	log.Println(names)
	sort.Strings(names)
	log.Println("Sorted:", names)

	log.Println(ages)
	sort.Ints(ages)
	log.Println("Sorted:", ages)

}

// Take pointer to type string and update using value at operator
func myFavColor(c *string) {
	*c = "Green"
}

// Take pointer to type struct and update without using value at operator
func myfavDeepColor(dC *deepColor) {
	dC.color = "Blue"
}

// Function exposed to outside packages
func MyColor() {

}
