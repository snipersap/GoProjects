# Type assertions in Go

When working with interfaces in Go, every once-in-awhile you'll need access to the underlying type of an interface value. You can cast an interface to its underlying type using a type assertion.

 A type assertion provides access to an interface value's underlying concrete value.
```go
t := i.(T)
```
This statement asserts that the interface value `i` holds the concrete type `T` and assigns the underlying `T` value to the variable `t`.

If `i` does not hold a `T`, the statement will trigger a panic. 

The example below shows how to safely access the radius field of s when s is an unknown type:

- we want to check if `s` is a circle in order to cast it into its underlying concrete type
- we know `s` is an instance of the `shape` interface, but we do not know if it's also a `circle`
- `c` is a new `circle` struct cast from `s`
- `ok` is true if `s` is indeed a `circle`, or false if `s` is NOT a `circle`
```go 
type shape interface {
	area() float64
}

type circle struct {
	radius float64
}

c, ok := s.(circle)
if !ok {
	// log an error if s isn't a circle
	log.Fatal("s is not a circle")
}

radius := c.radius
```

# Type Switches

A type switch makes it easy to do several type assertions in a series.

A type switch is similar to a regular switch statement, but the cases specify types instead of values.
```go 
func printNumericValue(num interface{}) {
	switch v := num.(type) {
	case int:
		fmt.Printf("%T\n", v)
	case string:
		fmt.Printf("%T\n", v)
	default:
		fmt.Printf("%T\n", v)
	}
}

func main() {
	printNumericValue(1)
	// prints "int"

	printNumericValue("1")
	// prints "string"

	printNumericValue(struct{}{})
	// prints "struct {}"
}

fmt.Printf("%T\n", v) prints the type of a variable.
```

# Assignment

Implement the getExpenseReport function.

- If the expense is an `email` then it should return the email's `toAddress` and the `cost` of the email.
- If the expense is an `sms` then it should return the sms's `toPhoneNumber` and its `cost`.
- If the expense has any other underlying type, just return an empty string and `0.0` for the cost.

Implement the getExpenseReportWithSwitch function and use type switches instead of if conditions. 

### Test results
```bash
---------------------------------
Inputs:     {isSubscribed:true body:Whoa there! toAddress:soldier@monty.com}
Expecting:  (soldier@monty.com, 0.11)
Actual:     (soldier@monty.com, 0.11)
Pass
---------------------------------
Inputs:     {isSubscribed:false body:Halt! Who goes there? toPhoneNumber:+155555509832}
Expecting:  (+155555509832, 2.1)
Actual:     (+155555509832, 2.1)
Pass
---------------------------------
Inputs:     {isSubscribed:false body:It is I, Arthur, son of Uther Pendragon, from the castle of Camelot. King of the Britons, defeator of the Saxons, sovereign of all England! toAddress:soldier@monty.com}
Expecting:  (soldier@monty.com, 6.95)
Actual:     (soldier@monty.com, 6.95)
Pass
---------------------------------
Inputs:     {isSubscribed:true body:Pull the other one! toAddress:arthur@monty.com}
Expecting:  (arthur@monty.com, 0.19)
Actual:     (arthur@monty.com, 0.19)
Pass
---------------------------------
Inputs:     {isSubscribed:true body:I am. And this my trusty servant Patsy. toPhoneNumber:+155555509832}
Expecting:  (+155555509832, 1.17)
Actual:     (+155555509832, 1.17)
Pass
---------------------------------
Inputs:     {}
Expecting:  (, 0)
Actual:     (, 0)
Pass
---------------------------------
Inputs:     {isSubscribed:true body:Whoa there! toAddress:soldier@monty.com}
Expecting:  (soldier@monty.com, 0.11)
Actual:     (soldier@monty.com, 0.11)
Pass
---------------------------------
Inputs:     {isSubscribed:false body:Halt! Who goes there? toPhoneNumber:+155555509832}
Expecting:  (+155555509832, 2.1)
Actual:     (+155555509832, 2.1)
Pass
---------------------------------
Inputs:     {isSubscribed:false body:It is I, Arthur, son of Uther Pendragon, from the castle of Camelot. King of the Britons, defeator of the Saxons, sovereign of all England! toAddress:soldier@monty.com}
Expecting:  (soldier@monty.com, 6.95)
Actual:     (soldier@monty.com, 6.95)
Pass
---------------------------------
Inputs:     {isSubscribed:true body:Pull the other one! toAddress:arthur@monty.com}
Expecting:  (arthur@monty.com, 0.19)
Actual:     (arthur@monty.com, 0.19)
Pass
---------------------------------
Inputs:     {isSubscribed:true body:I am. And this my trusty servant Patsy. toPhoneNumber:+155555509832}
Expecting:  (+155555509832, 1.17)
Actual:     (+155555509832, 1.17)
Pass
---------------------------------
Inputs:     {}
Expecting:  (, 0)
Actual:     (, 0)
Pass
---------------------------------
12 passed, 0 failed
--- PASS: Test (0.00s)
PASS
ok      type-assertion  0.541s
```