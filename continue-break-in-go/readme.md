# Continue & Break

Whenever we want to change the control flow of a loop we can use the continue and break keywords.
## continue

The continue keyword stops the current iteration of a loop and continues to the next iteration. continue is a powerful way to use the [guard clause](https://blog.boot.dev/clean-code/guard-clauses/) pattern within loops.
```go 
for i := 0; i < 10; i++ {
  if i % 2 == 0 {
    continue
  }
  fmt.Println(i)
}

// 1
// 3
// 5
// 7
// 9
```
## break

The break keyword stops the current iteration of a loop and exits the loop.
```go
for i := 0; i < 10; i++ {
  if i == 5 {
    break
  }
  fmt.Println(i)
}
// 0
// 1
// 2
// 3
// 4
```
## Assignment

As an Easter egg, we decided to reward our users with a free text message if they send a prime number of text messages this year.

Complete the `printPrimes` function. It should print all of the prime numbers up to and including max. It should skip any numbers that are not prime.

Here's the pseudocode:
```go 
printPrimes(max):
  for n in range(2, max+1):
    if n is 2:
      n is prime, print it
    if n is even:
      n is not prime, skip to next n
    for i in range (3, sqrt(n), 2):
      if i can be multiplied into n:
        n is not prime, skip to next n
    n is prime, print it
```
Easier logic implemented from [Efficient Approach - 1 Trial Division Method - O(√n) time and O(1) space](https://www.geeksforgeeks.org/check-for-prime-number/#efficient-approach-1-trial-division-method)
### Breakdown

    We skip even numbers because they can't be prime
    We only check up to the square root of n. A factor higher than the square root of n must multiply with a factor lower than the square root of n, meaning it has no chance of multiplying evenly into n.
        In your code, you can set the stop condition as i * i <= n
    We start checking at 2 because 1 is not prime

Note: This lesson is graded based on the output of your program, so don't leave any debugging print statements in your code.
#### Example Output

For the first test case, prime number up to 10:
```bash
Primes up to 10:
2
3
5
7
```
===============================================================

We only want you to print the numbers themselves, not the headings and delimiter.

## Additional exercise
Test Goldbach's Conjecture by printing the smallest number of primes that sum up to the given even number. 
Goldbach's conjecture: Every even number is the sum of 2 prime numbers. 

### Test results
```bash
Primes up to 10:
2
3
5
7
===============================================================
Primes up to 20:
2
3
5
7
11
13
17
19
===============================================================
Primes up to 29:
2
3
5
7
11
13
17
19
23
29
Smallest Primes adding up to 74:
(3,71)
Duration: 0
===============================================================
Smallest Primes adding up to 108:
(5,103)
Duration: 521900
===============================================================
Smallest Primes adding up to 584:
(7,577)
Duration: 0
===============================================================
Smallest Primes adding up to 998:
(7,991)
Duration: 525500
===============================================================
```