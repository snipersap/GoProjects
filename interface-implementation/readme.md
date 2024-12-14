# Interface Implementation

Interfaces are implemented implicitly.

A type never declares that it implements a given interface. If an interface exists and a type has the proper methods defined, then the type automatically fulfills that interface.
## Assignment

At Textio we have full-time employees and contract employees. We have been tasked with making a more general `employee` interface so that dealing with different employee types is simpler.

Run the code. You should see an error indicating the `contractor` type does not fulfill the `employee` interface.

Add the missing `getSalary` method to the `contractor` type so that it fulfills the `employee` interface.

A contractor's salary is their hourly pay multiplied by how many hours they work per year.

### Test results
```bash
$> go run main-go
Employee Donald Duck was paid $ 100000 this year!
Employee Mr. Scrooge was paid $ 344000 this year!
```
### RUN Tests
```bash
---------------------------------
Inputs:     {name:Bob salary:7300}
Expecting:  7300
Actual:     7300
Pass
---------------------------------
Inputs:     {name:Jill hourlyPay:872 hoursPerYear:982}
Expecting:  856304
Actual:     856304
Pass
---------------------------------
Inputs:     {name:Alice salary:10000}
Expecting:  10000
Actual:     10000
Pass
---------------------------------
Inputs:     {name:John hourlyPay:1000 hoursPerYear:1000}
Expecting:  1000000
Actual:     1000000
Pass
---------------------------------
4 passed, 0 failed
--- PASS: Test (0.00s)
```