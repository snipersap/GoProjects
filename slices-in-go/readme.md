# Slices in Go

99 times out of 100 you will use a `slice` instead of an `array` when working with ordered lists.

Arrays are fixed in size. Once you make an array like `[10]int` you can't add an 11th element.

A slice is a dynamically-sized, flexible view of the elements of an array.

[The zero value of slice is `nil`.](https://go.dev/tour/moretypes/12)

Non-nil slices always have an underlying array, though it isn't always specified explicitly. To explicitly create a slice on top of an array we can do:
```go 
primes := [6]int{2, 3, 5, 7, 11, 13}
mySlice := primes[1:4]
// mySlice = {3, 5, 7}
``` 
The syntax is:
```go 
arrayname[lowIndex:highIndex]
arrayname[lowIndex:]
arrayname[:highIndex]
arrayname[:]
```
Where `lowIndex` is inclusive and `highIndex` is exclusive.

`lowIndex`, `highIndex`, or both can be omitted to use the entire array on that side of the colon.

## Assignment

Retries are a premium feature now! Textio's free users only get 1 retry message, while pro members get an unlimited amount.

Complete the `getMessageWithRetriesForPlan` function. It takes a `plan` variable as input as well as an array of 3 messages. You've been provided with constants representing the `plan` types at the top of the file.

    If the plan is a pro plan, return all the strings from the messages input in a slice.
    If the plan is a free plan, return the first 2 strings from the messages input in a slice.
    If the plan isn't either of those, return a nil slice and an error that says unsupported plan.

### Test results
```bash
Test Passed:
Plan: free
Messages:
  - "Hello sir/madam can I interest you in a yacht?"
  - "Please I'll even give you an Amazon gift card?"
  - "You're missing out big time"
Expecting:
  - "Hello sir/madam can I interest you in a yacht?"
  - "Please I'll even give you an Amazon gift card?"
errString:
Actual:
  - "Hello sir/madam can I interest you in a yacht?"
  - "Please I'll even give you an Amazon gift card?"
errString:
Pass
---------------------------------
Test Passed:
Plan: pro
Messages:
  - "Hello sir/madam can I interest you in a yacht?"
  - "Please I'll even give you an Amazon gift card?"
  - "You're missing out big time"
Expecting:
  - "Hello sir/madam can I interest you in a yacht?"
  - "Please I'll even give you an Amazon gift card?"
  - "You're missing out big time"
errString:
Actual:
  - "Hello sir/madam can I interest you in a yacht?"
  - "Please I'll even give you an Amazon gift card?"
  - "You're missing out big time"
errString:
Pass
---------------------------------
Test Passed:
Plan: free
Messages:
  - "You can get a good look at a T-bone by sticking your head up a bull's ass, but wouldn't you rather take the butcher's word for it?"
  - "Wouldn't you?"
  - "Wouldn't you???"
Expecting:
  - "You can get a good look at a T-bone by sticking your head up a bull's ass, but wouldn't you rather take the butcher's word for it?"
  - "Wouldn't you?"
errString:
Actual:
  - "You can get a good look at a T-bone by sticking your head up a bull's ass, but wouldn't you rather take the butcher's word for it?"
  - "Wouldn't you?"
errString:
Pass
---------------------------------
Test Passed:
Plan: pro
Messages:
  - "You can get a good look at a T-bone by sticking your head up a bull's ass, but wouldn't you rather take the butcher's word for it?"
  - "Wouldn't you?"
  - "Wouldn't you???"
Expecting:
  - "You can get a good look at a T-bone by sticking your head up a bull's ass, but wouldn't you rather take the butcher's word for it?"
  - "Wouldn't you?"
  - "Wouldn't you???"
errString:
Actual:
  - "You can get a good look at a T-bone by sticking your head up a bull's ass, but wouldn't you rather take the butcher's word for it?"
  - "Wouldn't you?"
  - "Wouldn't you???"
errString:
Pass
---------------------------------
Test Passed:
Plan: invalid plan
Messages:
  - "You can get a good look at a T-bone by sticking your head up a bull's ass, but wouldn't you rather take the butcher's word for it?"
  - "Wouldn't you?"
  - "Wouldn't you???"
Expecting:
  <nil>
errString:  unsupported plan
Actual:
  []
errString:  unsupported plan
Pass
---------------------------------
5 passed, 0 failed
PASS
ok      slices-in-go    0.142s
```
