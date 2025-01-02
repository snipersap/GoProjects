# Formatting strings review

A convenient way to format strings in Go is by using the standard library's fmt.Sprintf() function. It's a string interpolation function, similar to Python's f-strings. The %v substring uses the type's default formatting, which is often what you want.
Default values
```go
const name = "Kim"
const age = 22
s := fmt.Sprintf("%v is %v years old.", name, age)
// s = "Kim is 22 years old."
```
The equivalent Python code:
```python
name = "Kim"
age = 22
s = f"{name} is {age} years old."
# s = "Kim is 22 years old."
```
Rounding floats
```go 
fmt.Printf("I am %f years old", 10.523)
// I am 10.523000 years old

// The ".2" rounds the number to 2 decimal places
fmt.Printf("I am %.2f years old", 10.523)
// I am 10.52 years old
```
## Assignment

We need better error logs for our backend developers to help them debug their code.

Complete the `getSMSErrorString()` function. It should return a string with this format:

    SMS that costs $COST to be sent to 'RECIPIENT' can not be sent

`COST` is the cost of the SMS, always showing the price formatted to 2 decimal places.
`RECIPIENT` is the stringified representation of the recipient's phone number

Be sure to include the $ symbol and the single quotes

### Test results
```bash
=== RUN   Test
---------------------------------
Inputs:     (1.4, +1 (435) 555 0923)
Expecting:  SMS that costs $1.40 to be sent to '+1 (435) 555 0923' can not be sent
Actual:     SMS that costs $1.40 to be sent to '+1 (435) 555 0923' can not be sent
Pass
---------------------------------
Inputs:     (2.1, +2 (702) 555 3452)
Expecting:  SMS that costs $2.10 to be sent to '+2 (702) 555 3452' can not be sent
Actual:     SMS that costs $2.10 to be sent to '+2 (702) 555 3452' can not be sent
Pass
---------------------------------
Inputs:     (32.1, +1 (801) 555 7456)
Expecting:  SMS that costs $32.10 to be sent to '+1 (801) 555 7456' can not be sent
Actual:     SMS that costs $32.10 to be sent to '+1 (801) 555 7456' can not be sent
Pass
---------------------------------
Inputs:     (14.4, +1 (234) 555 6545)
Expecting:  SMS that costs $14.40 to be sent to '+1 (234) 555 6545' can not be sent
Actual:     SMS that costs $14.40 to be sent to '+1 (234) 555 6545' can not be sent
Pass
---------------------------------
4 passed, 0 failed
--- PASS: Test (0.00s)
PASS
ok      error-intf-format-strings       0.245s
```