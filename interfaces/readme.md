# Interfaces in Go

[Interfaces](https://go.dev/tour/methods/9) are just collections of method signatures. A type "implements" an interface if it has methods that match the interface's method signatures.

In the following example, a "shape" must be able to return its area and perimeter. Both `rect` and `circle` fulfill the interface.
```go
type shape interface {
  area() float64
  perimeter() float64
}

type rect struct {
    width, height float64
}
func (r rect) area() float64 {
    return r.width * r.height
}
func (r rect) perimeter() float64 {
    return 2*r.width + 2*r.height
}

type circle struct {
    radius float64
}
func (c circle) area() float64 {
    return math.Pi * c.radius * c.radius
}
func (c circle) perimeter() float64 {
    return 2 * math.Pi * c.radius
}
```
When a type implements an interface, it can then be used as that interface type.
```go
func printShapeData(s shape) {
	fmt.Printf("Index: %v - Area: %v - Perimeter: %v\n", i, s.area(), s.perimeter())
}
```
Because we say the input is of type shape, we know that any argument must implement the .area() and .perimeter() methods.

As an example, because the [empty interface](https://go.dev/tour/methods/14) doesn't require a type to have any methods at all, every type automatically implements the empty interface, written as:
```go
interface{}
```
Interfaces allow you to focus on what a type does rather than how it's built. They can help you write more flexible and reusable code by defining behaviors (like methods) that different types can share. This makes it easy to swap out or update parts of your code without changing everything else.
## Assignment

The `birthdayMessage` and `sendingReport` structs already have implementations of the `getMessage` method. The `getMessage` method returns a string, and any type that implements the method can be considered a message (meaning it implements the `message` interface).

    Add the getMessage() method signature as a requirement on the message interface.
    Complete the sendMessage function. It should return:
        The content of the message.
        The cost of the message, which is the length of the message multiplied by 3.

_Notice that your code doesn't care at all about whether a specific message is a `birthdayMessage` or a `sendingReport`!_
### Tip

The length of a string can be obtained using the len function, which [returns the number of bytes](https://pkg.go.dev/builtin#len).
```go
s := "Hello, World!"
fmt.Println(len(s))
// 13
```

### Test results
```bash
$> go run main.go
Message: Hi Donald Duck, it is your birthday on 2024-12-14T22:51:49+01:00 , and it cost you (in cents): 192
Message: Your "Communication insights" report is ready. You've sent 400 messages. , and it cost you (in cents): 216
```
### === RUN Test
```bash
$> go test
Message: Hi John Doe, it is your birthday on 1994-03-21T00:00:00Z , and it cost you (in cents): 168
---------------------------------
Inputs:     {birthdayTime:{wall:0 ext:62899804800 loc:<nil>} recipientName:John Doe}
Expecting:  (Hi John Doe, it is your birthday on 1994-03-21T00:00:00Z, 168)
Actual:     (Hi John Doe, it is your birthday on 1994-03-21T00:00:00Z, 168)
Pass
Message: Your "First Report" report is ready. You've sent 10 messages. , and it cost you (in cents): 183
---------------------------------
Inputs:     {reportName:First Report numberOfSends:10}
Expecting:  (Your "First Report" report is ready. You've sent 10 messages., 183)
Actual:     (Your "First Report" report is ready. You've sent 10 messages., 183)
Pass
Message: Hi Bill Deer, it is your birthday on 1934-05-01T00:00:00Z , and it cost you (in cents): 171
---------------------------------
Inputs:     {birthdayTime:{wall:0 ext:61009891200 loc:<nil>} recipientName:Bill Deer}
Expecting:  (Hi Bill Deer, it is your birthday on 1934-05-01T00:00:00Z, 171)
Actual:     (Hi Bill Deer, it is your birthday on 1934-05-01T00:00:00Z, 171)
Pass
Message: Your "Second Report" report is ready. You've sent 20 messages. , and it cost you (in cents): 186
---------------------------------
Inputs:     {reportName:Second Report numberOfSends:20}
Expecting:  (Your "Second Report" report is ready. You've sent 20 messages., 186)
Actual:     (Your "Second Report" report is ready. You've sent 20 messages., 186)
Pass
---------------------------------
4 passed, 0 failed
--- PASS: Test (0.00s)
PASS
ok      interfaces      0.505s
```