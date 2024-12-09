# Anonymous Structs in Go

An anonymous struct is just like a normal struct, but it is defined without a name and therefore cannot be referenced elsewhere in the code.

To create an anonymous struct, just instantiate the instance immediately using a second pair of brackets after declaring the type:

myCar := struct {
  brand string
  model string
} {
  brand: "tesla",
  model: "model 3",
}

You can even nest anonymous structs as fields within other structs:

type car struct {
  brand string
  model string
  doors int
  mileage int
  // wheel is a field containing an anonymous struct
  wheel struct {
    radius int
    material string
  }
}

When should you use an anonymous struct?

In general, prefer named structs. Named structs make it easier to read and understand your code, and they have the nice side-effect of being reusable. I sometimes use anonymous structs when I know I won't ever need to use a struct again. For example, sometimes I'll use one to create the shape of some JSON data in HTTP handlers.

If a struct is only meant to be used once, then it makes sense to declare it in such a way that developers down the road won’t be tempted to accidentally use it again.

You can read more about [anonymous structs here](https://blog.boot.dev/golang/anonymous-structs-golang/) if you're curious.

### Test results
```bash
$> go run .\main.go
My car is a bmw and is named little johhny
My Second Car is a jaguar and has a wheel of raidius 6.2 and its wheel size is S
```
