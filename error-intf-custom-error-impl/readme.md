# The Error Interface

Because errors are just interfaces, you can build your own custom types that implement the error interface. Here's an example of a userError struct that implements the error interface:
```go
type userError struct {
    name string
}

func (e userError) Error() string {
    return fmt.Sprintf("%v has a problem with their account", e.name)
}
```
It can then be used as an error:
```go 
func sendSMS(msg, userName string) error {
    if !canSendToUser(userName) {
        return userError{name: userName}
    }
    ...
}
```
## Assignment

Our users frequently try to run custom analytics queries on their message deliverability metrics, and end up writing a bad query that tries to divide a number by zero. It's become such a problem that we need to make a new type of error for division by zero.

Update the code so that the `divideError` type implements the error interface. Its `Error()` method should just return a string formatted in the following way:

    can not divide DIVIDEND by zero

Where `DIVIDEND` is the actual dividend of the `divideError`. Use the `%v` verb to format the dividend as a `float`.

### Test results
```bash
=== RUN   TestDivide
---------------------------------
Inputs:     (10, 2)
Expecting:  (5, )
Actual:     (5, )
Pass
---------------------------------
Inputs:     (15, 3)
Expecting:  (5, )
Actual:     (5, )
Pass
---------------------------------
Inputs:     (10, 0)
Expecting:  (0, can not divide 10 by zero)
Actual:     (0, can not divide 10 by zero)
Pass
---------------------------------
Inputs:     (15, 0)
Expecting:  (0, can not divide 15 by zero)
Actual:     (0, can not divide 15 by zero)
Pass
---------------------------------
Inputs:     (100, 10)
Expecting:  (10, )
Actual:     (10, )
Pass
---------------------------------
Inputs:     (16, 4)
Expecting:  (4, )
Actual:     (4, )
Pass
---------------------------------
Inputs:     (30, 6)
Expecting:  (5, )
Actual:     (5, )
Pass
---------------------------------
7 passed, 0 failed
--- PASS: TestDivide (0.00s)
PASS
ok      error-intf-custom-error-impl    0.274s
```