package main

import (
	"errors"
	"fmt"
	"log"
)

func main() {

	x := float32(10.0)
	y := float32(0.0)
	res, err := divide(x, y)
	if err != nil {
		log.Println("Divide error:", err.Error())
	} else {
		fmt.Printf("Result of dividing %f by %f is: %f", x, y, res)
	}
}

func divide(x float32, y float32) (float32, error) {

	if y == 0 {
		return 0.0, errors.New("divide by zero error")
	} else {
		return x / y, nil
	}
}
