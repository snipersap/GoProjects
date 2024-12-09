# Embedded Structs

Go is not an object-oriented language. However, embedded structs provide a kind of data-only inheritance that can be useful at times. Keep in mind, Go doesn't support classes or inheritance in the complete sense, but embedded structs are a way to elevate and share fields between struct definitions.

```go
type car struct {
  brand string
  model string
}

type truck struct {
  // "car" is embedded, so the definition of a
  // "truck" now also additionally contains all
  // of the fields of the car struct
  car
  bedSize int
}
```
## Embedded vs nested

Unlike nested structs, an embedded struct's fields are accessed at the top level like normal fields.
Like nested structs, you assign the promoted fields with the embedded struct in a composite literal.

```go
lanesTruck := truck{
  bedSize: 10,
  car: car{
    brand: "toyota",
    model: "camry",
  },
}

fmt.Println(lanesTruck.bedSize)

// embedded fields promoted to the top-level
// instead of lanesTruck.car.brand
fmt.Println(lanesTruck.brand)
fmt.Println(lanesTruck.model)
```
## Assignment

At Textio, a "user" struct represents an account holder, and a "sender" is just a "user" with some additional "sender" specific data. A "sender" is a user that has a rateLimit field that tells us how many messages they are allowed to send.

Fix the bug by embedding the proper struct in the other.

### Test results
```bash
$> go test
---------------------------------
Inputs:     (10000, Deborah, 18055558790)
Expecting:
====================================
Sender name: Deborah
Sender number: 18055558790
Sender rateLimit: 10000
====================================

Actual:
====================================
Sender name: Deborah
Sender number: 18055558790
Sender rateLimit: 10000
====================================

Pass
---------------------------------
Inputs:     (5000, Jason, 18055558791)
Expecting:
====================================
Sender name: Jason
Sender number: 18055558791
Sender rateLimit: 5000
====================================

Actual:
====================================
Sender name: Jason
Sender number: 18055558791
Sender rateLimit: 5000
====================================

Pass
---------------------------------
Inputs:     (1000, Jill, 18055558792)
Expecting:
====================================
Sender name: Jill
Sender number: 18055558792
Sender rateLimit: 1000
====================================

Actual:
====================================
Sender name: Jill
Sender number: 18055558792
Sender rateLimit: 1000
====================================

Pass
---------------------------------
3 passed, 0 failed
PASS
ok      embedded-struct 0.131s
```