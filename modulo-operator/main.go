package main

import (
	"fmt"
)

func fizzbuzz() {
	//Print 1 to 100 in each line
	max := 100
	for i := 1; i <= max; i++ {
		fizzbuzz := 0
		if i%3 == 0 {
			fizzbuzz = 3
		}
		if i%5 == 0 {
			fizzbuzz += 2
		}
		switch fizzbuzz {
		case 2:
			fmt.Println("buzz")
		case 3:
			fmt.Println("fizz")
		case 5:
			fmt.Println("fizzbuzz")
		default:
			fmt.Printf("%d\n", i)
		}
	}
}

// don't touch below this line

func main() {
	fizzbuzz()
}
