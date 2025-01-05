# Loops in Go

The basic loop in Go is written in standard C-like syntax:
```go 
for INITIAL; CONDITION; AFTER{
  // do something
}
```
`INITIAL` is run once at the beginning of the loop and can create
variables within the scope of the loop.

`CONDITION` is checked before each iteration. If the condition doesn't pass
then the loop breaks.

`AFTER` is run after each iteration.

For example:
```go
for i := 0; i < 10; i++ {
  fmt.Println(i)
}
// Prints 0 through 9
```
## Assignment

At Textio we have a dynamic formula for determining how much a batch of bulk messages costs to send. Complete the `bulkSend()` function.

It should return the total cost (as a `float64`) to send a batch of `numMessages` messages. Each message costs 1.0, plus an additional fee. The fee structure is:

- 1st message: 1.0 + 0.00
- 2nd message: 1.0 + 0.01
- 3rd message: 1.0 + 0.02
- 4th message: 1.0 + 0.03
- ...

Use a loop to calculate the total cost and return it.

### Test results
```bash
=== RUN   Test
---------------------------------
Inputs:     (10)
Expecting:  10.45
Actual:     10.45
Pass
---------------------------------
Inputs:     (20)
Expecting:  21.90
Actual:     21.90
Pass
---------------------------------
Inputs:     (0)
Expecting:  0.00
Actual:     0.00
Pass
---------------------------------
Inputs:     (1)
Expecting:  1.00
Actual:     1.00
Pass
---------------------------------
Inputs:     (5)
Expecting:  5.10
Actual:     5.10
Pass
---------------------------------
Inputs:     (30)
Expecting:  34.35
Actual:     34.35
Pass
---------------------------------
6 passed, 0 failed
--- PASS: Test (0.00s)
PASS
ok      loops-in-Textio 0.114s
```