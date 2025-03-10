package main

import (
	"fmt"
	"time"
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

func printPrimeCombo(even int) {
	for i := 2; i < even/2; i++ {
		if isPrime(i) && isPrime(even-i) {
			fmt.Printf("(%v,%v)\n", i, even-i)
			break
		}
	}
}

func isPrime(num int) bool {
	if num <= 1 {
		return false
	}
	for i := 2; i*i <= num; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true

}

// don't edit below this line

func test(max int) {
	fmt.Printf("Primes up to %v:\n", max)
	printPrimes(max)
	fmt.Println("===============================================================")
}

func testGoldbachsConjecture(even int) {
	fmt.Printf("Smallest Primes adding up to %v:\n", even)
	start := time.Now()
	printPrimeCombo(even)
	fmt.Println("Duration:", time.Since(start).Nanoseconds())
	fmt.Println("===============================================================")
}

func main() {
	test(10)
	test(20)
	test(29)
	testGoldbachsConjecture(74)
	testGoldbachsConjecture(108)
	testGoldbachsConjecture(584)
	testGoldbachsConjecture(998)
}
