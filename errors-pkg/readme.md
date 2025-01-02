# The Errors Package

The Go standard library provides an "errors" package that makes it easy to deal with errors.

Read the godoc for the errors.New() function, but here's a simple example:
```go
var err error = errors.New("something went wrong")
```
## Assignment

Textio's software architects may have overcomplicated the requirements from the last coding assignment... oops. All we needed was a new generic error message that returns the string no dividing by 0 when a user attempts to get us to perform the taboo.

Complete the divide function. Use the errors.New() function to return an error when y == 0 that reads "no dividing by 0".

### Test results
```bash
=== RUN   TestDivide
---------------------------------
Inputs:     (10, 0)
Expecting:  (0, no dividing by 0)
Actual:     (0, no dividing by 0)
Pass
---------------------------------
Inputs:     (10, 2)
Expecting:  (5, )
Actual:     (5, )
Pass
---------------------------------
Inputs:     (15, 30)
Expecting:  (0.5, )
Actual:     (0.5, )
Pass
---------------------------------
Inputs:     (6, 3)
Expecting:  (2, )
Actual:     (2, )
Pass
---------------------------------
Inputs:     (0, 10)
Expecting:  (0, )
Actual:     (0, )
Pass
---------------------------------
Inputs:     (100, 0)
Expecting:  (0, no dividing by 0)
Actual:     (0, no dividing by 0)
Pass
---------------------------------
Inputs:     (-10, -2)
Expecting:  (5, )
Actual:     (5, )
Pass
---------------------------------
Inputs:     (-10, 2)
Expecting:  (-5, )
Actual:     (-5, )
Pass
---------------------------------
8 passed, 0 failed
--- PASS: TestDivide (0.00s)
PASS
ok      errors-pkg      0.416s
Running tool: C:\Program Files\Go\bin\go.exe test -timeout 30s -run ^TestDivide$ errors-pkg
```