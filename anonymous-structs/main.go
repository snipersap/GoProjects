package main

import (
	"fmt"
)

type car struct {
	brand string
	wheel struct {
		radius float64
		size   string
	}
}

func main() {
	myCar := struct {
		brand string
		name  string
	}{
		brand: "bmw",
		name:  "little johhny",
	}
	fmt.Println("My car is a", myCar.brand, "and is named", myCar.name)

	var mySecondCar car
	mySecondCar.brand = "jaguar"
	mySecondCar.wheel.radius = 6.2
	mySecondCar.wheel.size = "S"
	fmt.Println("My Second Car is a", mySecondCar.brand, "and has a wheel of raidius", mySecondCar.wheel.radius, "and its wheel size is", mySecondCar.wheel.size)

}
