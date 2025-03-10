package main

import (
	"fmt"
)

func printPrimes(max int) {
	for n := range max + 1 {
		if n <= 1 {
			continue
		}
		//logic: https://www.geeksforgeeks.org/check-for-prime-number/
		prime := true
		for i := 2; i*i <= n; i++ {
			if n%i == 0 {
				prime = false
				break
			}
		}
		if prime {
			fmt.Println(n)
		}
	}
}

// don't edit below this line

func test(max int) {
	fmt.Printf("Primes up to %v:\n", max)
	printPrimes(max)
	fmt.Println("===============================================================")
}

func main() {
	test(10)
	test(20)
	test(29)
}
