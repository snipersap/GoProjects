# Nested structs in Go

Structs can be nested to represent more complex entities:

type car struct {
  brand string
  model string
  doors int
  mileage int
  frontWheel wheel
  backWheel wheel
}

type wheel struct {
  radius int
  material string
}

The fields of a struct can be accessed using the dot . operator.

myCar := car{}
myCar.frontWheel.radius = 5

## Assignment

Textio has a bug, we've been sending texts that are missing critical bits of information! Before we send text messages in Textio, we must check to make sure the required fields have non-zero values.

Notice that the user struct is a nested struct within the messageToSend struct. Both sender and recipient are user struct types.

Complete the canSendMessage function. It should return true only if the sender and recipient fields each contain a name and a number. If any of the default zero values are present, return false instead.
