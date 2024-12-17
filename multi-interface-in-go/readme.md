# Multiple Interfaces

A type can implement any number of interfaces in Go. For example, the empty interface, interface{}, is always implemented by every type because it has no requirements.
## Assignment

Complete the required methods so that the email type implements both the `expense` and `formatter` interfaces.
### cost()

If the email is not "subscribed", then the cost is 5 cents for each character in the body. If it is, then the cost is 2 cents per character.

Return the total cost of the entire email in cents.
### format()

The format method should return a string in this format:
```bash
'CONTENT' | Subscribed
```
If the email is not subscribed, change the second part to "Not Subscribed":
```bash
'CONTENT' | Not Subscribed
```
The single quotes are included in the string, and CONTENT is the email's body. For example:
```bash
'Hello, World!' | Subscribed
```
Note: you may want to import the `fmt` package and use `Sprintf`.

### Test Results
```bash
---------------------------------
Inputs:     (hello there, true)
Expecting:  (22, 'hello there' | Subscribed)
Actual:     (22, 'hello there' | Subscribed)
Pass
---------------------------------
Inputs:     (general kenobi, false)
Expecting:  (70, 'general kenobi' | Not Subscribed)
Actual:     (70, 'general kenobi' | Not Subscribed)
Pass
---------------------------------
Inputs:     (i hate sand, true)
Expecting:  (22, 'i hate sand' | Subscribed)
Actual:     (22, 'i hate sand' | Subscribed)
Pass
---------------------------------
Inputs:     (it's coarse and rough and irritating, false)
Expecting:  (180, 'it's coarse and rough and irritating' | Not Subscribed)
Actual:     (180, 'it's coarse and rough and irritating' | Not Subscribed)
Pass
---------------------------------
Inputs:     (and it gets everywhere, true)
Expecting:  (44, 'and it gets everywhere' | Subscribed)
Actual:     (44, 'and it gets everywhere' | Subscribed)
Pass
---------------------------------
5 passed, 0 failed
--- PASS: Test (0.00s)
PASS
ok      multi-interface-in-go   0.432s
```