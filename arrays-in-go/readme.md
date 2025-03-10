# Arrays in Go

Arrays are fixed-size groups of variables of the same type. For example, [4]string is an array of 4 values of type `string`.

To declare an array of 10 integers:
```go 
var myInts [10]int
```
or to declare an initialized literal:
```go 
primes := [6]int{2, 3, 5, 7, 11, 13}
```
## Assignment

When our clients don't respond to a message, they can be reminded with up to 2 additional messages.

Complete the `getMessageWithRetries` function. It takes three strings and returns:

    An array of 3 strings
    An array of 3 integers

The returned string array contains the original messages. The first is the primary message, the second is the first reminder, and the third is the last reminder.

The integers in the integer array represent the cost of sending each message. The cost of each message is equal to the length of the message, plus the length of any previous messages. For example:

    "hello" costs 5
    "world" costs 5, adding "hello" makes total cost 10 (5 + 5)
    "!" costs 1, adding the previous messages makes total cost 11 (5 + 5 + 1)

### Test results
```bash
Test Passed:
Inputs:
  - Hello sir/madam can I interest you in a yacht?
  - Please I'll even give you an Amazon gift card?
  - You're missing out big time
Expecting:
  - Hello sir/madam can I interest you in a yacht?
  - Please I'll even give you an Amazon gift card?
  - You're missing out big time
[46 92 119]
Actual:
  - Hello sir/madam can I interest you in a yacht?
  - Please I'll even give you an Amazon gift card?
  - You're missing out big time
[46 92 119]
Pass
---------------------------------
Test Passed:
Inputs:
  - It's the spring fling sale!
  - Don't miss this event!
  - Last chance.
Expecting:
  - It's the spring fling sale!
  - Don't miss this event!
  - Last chance.
[27 49 61]
Actual:
  - It's the spring fling sale!
  - Don't miss this event!
  - Last chance.
[27 49 61]
Pass
---------------------------------
Test Passed:
Inputs:
  - Put that coffee down!
  - Coffee is for closers
  - Always be closing
Expecting:
  - Put that coffee down!
  - Coffee is for closers
  - Always be closing
[21 42 59]
Actual:
  - Put that coffee down!
  - Coffee is for closers
  - Always be closing
[21 42 59]
Pass
---------------------------------
3 passed, 0 failed
PASS
ok      arrays-in-go    0.381s
```
