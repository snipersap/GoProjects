Structs in Go

We use structs in Go to represent structured data. It's often convenient to group different types of variables together. For example, if we want to represent a car we could do the following:

```go
type car struct {
	brand      string
	model      string
	doors      int
	mileage    int
	frontWheel wheel
	backWheel  wheel
}
```
This creates a new struct type called car. All cars have a brand, model, doors and mileage.

Structs in Go are often used to represent data that you might use a dictionary or object for in other languages.
## Assignment

Complete the definition of the messageToSend struct. It needs two fields:

phoneNumber - an integer
message - a string.

### Test results
```bash
$> go run .\main.go
This is my Phone Number: 12346

$> go test
---------------------------------
Inputs:     (148255510981, Thanks for signing up)
Expecting:  Sending message: 'Thanks for signing up' to: 148255510981
Actual:     Sending message: 'Thanks for signing up' to: 148255510981
Pass
---------------------------------
Inputs:     (148255510982, Love to have you aboard!)
Expecting:  Sending message: 'Love to have you aboard!' to: 148255510982
Actual:     Sending message: 'Love to have you aboard!' to: 148255510982
Pass
---------------------------------
Inputs:     (148255510983, We're so excited to have you)
Expecting:  Sending message: 'We're so excited to have you' to: 148255510983
Actual:     Sending message: 'We're so excited to have you' to: 148255510983
Pass
---------------------------------
Inputs:     (148255510984, )
Expecting:  Sending message: '' to: 148255510984
Actual:     Sending message: '' to: 148255510984
Pass
---------------------------------
Inputs:     (148255510985, Hello, World!)
Expecting:  Sending message: 'Hello, World!' to: 148255510985
Actual:     Sending message: 'Hello, World!' to: 148255510985
Pass
---------------------------------
5 passed, 0 failed
PASS
ok      struct-in-go    0.109s
```
