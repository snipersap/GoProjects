package main

import (
	"encoding/hex"
	"errors"
	"fmt"
	"math"
	"reflect"
	"strconv"
	"unicode/utf8"
)

func main() {

	//1. Printf
	fmt.Printf("Printf: Hello there %s!\n", "Ashok")
	//2. Sprintf round to 2 decimal places
	name := fmt.Sprintf("Sprintf: Hello %s from Sprintf, you are %.2f years old", "Priya", 10.456)
	fmt.Println(name)

	//3. Strconv package
	i, _ := strconv.Atoi("-52")
	fmt.Printf("Converting -52 from string to integer gives us:%d, which is of type %T\n", i, i)

	s := strconv.Itoa(-256)
	fmt.Printf("Converting -256 from integer to string gives us:%s, which is of type %s\n", s, reflect.TypeOf(s))

	b, _ := strconv.ParseBool("true")
	fmt.Printf("Parsed boolean value of true = %t\n", b)
	b1, _ := strconv.ParseBool("hello")
	fmt.Printf("Parsed boolean value of hello = %t\n", b1)
	b2, _ := strconv.ParseBool("false")
	fmt.Printf("Parsed boolean value of false = %t\n", b2)

	f, _ := strconv.ParseFloat("-2673547.36584685", 32) //TODO: Why is the precision limited to .25 only?
	f1, _ := strconv.ParseFloat("-08008111.47237474274725742", 64)
	fmt.Printf("Parse string -2673547.36584685 to float32 %.9f, with type %T\n", f, f)
	fmt.Printf("Parse string -08008111.47237474274725742 to float64 %5.9f, with type %T\n", f1, f1)

	s1 := strconv.FormatFloat(-256.2458777, 'e', 6, 64)
	fmt.Printf("Converting floating point number -256.2458777 to string %s, of type %T\n", s1, s1)

	//4. Emojis and runes
	r := "🐻"
	var r1 rune = '🐻'
	fmt.Printf("The bear emoji has length %d bytes and type %T, and length of runes %d\n", len(r), r, utf8.RuneCountInString(r))
	fmt.Printf("The value of rune r1 is %v in unicode, while as a rune it is %c\n", r1, r1)

	//5. Casting and truncating variables
	//Float to int conversion
	ageF := 2.6
	ageI := int(ageF)
	fmt.Printf("Although ageF,a floating point number, is %f, when stored in an integer ageI, it becomes %d\n", ageF, ageI)
	//Int to String conversion
	pass := 12345
	passStr := string(pass)
	fmt.Printf("Converting password from int to string gives:%s, of type %T, which means direct type casting to string outputs a string of one rune and not digits\n", passStr, passStr)
	password := strconv.Itoa(12345)
	fmt.Printf("Using Itoa function of package strconv, we get the password %s, as type %T\n", password, password)

	//6. Same line declarations
	brother, age := "Mark", 26
	fmt.Printf("Same line declarations, Brother's name is %s and age is %d\n", brother, age)

	//7. Declaring constants
	const pi float64 = 3.1456
	const daysInYear int = 365
	const daysInLeapYear string = "366"
	fmt.Printf("Constants declared were pi:%f, days in a year: %d, and days in a leap year:%s\n", pi, daysInYear, daysInLeapYear)

	//8. Computed constants
	const daysNormalYear int = 365
	const day int = 1
	const daysInLeapYears = daysNormalYear + day
	fmt.Printf("Days in a leap year, created and computed from constants = %d\n", daysInLeapYears)

	//9. Conditionals
	fmt.Println("Conditionals: ")
	height, length, breadth := 100, 200, 300
	if height > 100 {
		fmt.Println("It is quite high!")
	} else if length > 200 {
		fmt.Println("It is quite long!")
	} else {
		fmt.Println("Neither high nor long!")
	}
	if breadth > 200 {
		fmt.Println("The breadth is adequate.")
	}

	//10. Conditionals with initial statement in IF
	fmt.Println("Conditionals with initial statement in IF:")
	if age := 68; age > 67 {
		fmt.Println("Time for retirement! Age more than 67. Good for limiting scope of age to the IF condition.")
	}

	//11. Conditionals with Switch
	fmt.Println("Conditionals with Switch:")
	switch version := 2.1; version {
	case 0:
		fmt.Println("Intial version.")
	case 1:
		fmt.Println("First Release.")
	case 1.1:
		fmt.Println("First patch of first release.")
	case 2.0:
		fmt.Println("Second Release.")
	case 2.1:
		fmt.Println("First patch of Second Release.")
	default:
		fmt.Println("We released more than 2 releases. ")
	}

	fmt.Println("Switch fallthrough:")
	switch ver := 1.4; ver {
	case 1.0:
		fallthrough
	case 1.1:
		fallthrough
	case 1.2:
		fallthrough
	case 1.3:
		fallthrough
	case 1.4:
		fallthrough
	case 1.5:
		fmt.Println("First Release and/or its patches.")
	}

	//12. Functions
	fmt.Println("Simple Functions: and pass by value")
	a1, a2 := 1, 2
	sum, _ := add(a1, a2)
	fmt.Printf("Before adding, a1:%d, a2:%d.\n", a1, a2)
	fmt.Printf("Sum from the add function is %d.\n", sum)
	fmt.Printf("After adding, a1:%d, a2:%d, have not changed as add() args were passed with value.\n", a1, a2)

	//13. Naked returns
	x, y := getLocation()
	fmt.Printf("Get coordinates with naked return: X:%f, y:%f\n", x, y)

	//14. Early Return
	fmt.Println("Early returns with age:")
	fmt.Printf("The person is (a/an) %s.\n", earlyReturn(56))

	//15. Functions as arguments to another function
	fmt.Println("Functions as arguments: multiplication of 3 numbers")
	fmt.Println("Multiplying 2,3 and 4:", arithmetic(2, 3, 4, mul))

	fmt.Println("Functions as arguments: Subtraction of 3 numbers")
	fmt.Println("Subtraction of 2,3 and 4 in that sequence:", arithmetic(2, 3, 4, sub))

	//16. Anonymous Functions
	fmt.Println("Addition via anonymous function:")
	sumOfAdd := func(a, b int) int {
		return a + b
	}(2, 4)
	fmt.Println("The sum of 2 numbers through an anonymous function is:", sumOfAdd)

	//17. Defer function
	fmt.Println("Defer functions:")
	completeAction(7)

	//18. Closures
	fmt.Println("Closures access variables outside its own function body.")
	writeStory := baseline()
	writeStory("Little Red riding Hood")
	writeStory("Went through the jungle to visit her Granny")
	writeStory("The bad wolf waited for his prey disguised as her Granny")
	writeStory("Little Red Riding Hood was too smart for the wolf")
	fmt.Println(writeStory("She got away safely"))

	//19. Structs - Nested, Named, anonymous and embedded
	//Definition
	type seat struct {
		upholstery string
		color      string
	}

	type light struct {
		kind         string
		power        string
		motionSensor bool
	}

	type car struct {
		doors     int
		brand     string
		color     string
		model     string
		frontSeat seat //nested
		rearSeat  seat //nested
		wheel     struct {
			kind     string
			pressure float64
		} //anonymous nested
		light //embedded
	}
	//Initialization
	frontSeat := seat{
		upholstery: "leather",
		color:      "#abcde",
	}

	myCar := car{
		brand:     "BMW",
		color:     "#fff888",
		model:     "i20FR",
		doors:     5,
		frontSeat: frontSeat, //with variable
		wheel: struct {
			kind     string
			pressure float64
		}{pressure: 3.4}, //direct
		light: light{kind: "LED", power: "3000 Lumens"}, //direct
	}
	//access nested struct
	myCar.rearSeat.upholstery = "polymer"
	myCar.rearSeat.color = "#zzzhhr"
	//access field with anonymous nested struct
	myCar.wheel.kind = "M+S"
	//access embedded struct
	myCar.motionSensor = true
	//Usage
	fmt.Println("My car configuration via nested, named and anonymous struct is:", myCar)

	//20. Hexadecimal values
	h1 := 0x1a
	fmt.Printf("Hex value is: %x, Decimal value is %d\n", h1, h1)
	d2 := 42
	h2 := fmt.Sprintf("%#x", d2)
	fmt.Printf("Hex value for 42 is %v\n", h2)
	//Encode []bytes to hex
	bt := []byte("Golang")
	h3 := hex.EncodeToString(bt)
	fmt.Println("Encoded GoLang string in hex is:", h3)
	btFromHex, _ := hex.DecodeString(h3)
	fmt.Println("Decoded back to String is:", string(btFromHex))

	//21. Receivers of Struct
	r2 := rect{length: 10, breadth: 20}
	fmt.Println("Area of the rectangle via receiver of struct is:", r2.area())

	//22. Optimize size of Struct
	typ := reflect.TypeOf(myCar)
	fmt.Printf("The size of the type of myCar (a struct) is:%d bytes\n", typ.Size())

	typ = reflect.TypeOf(r2)
	fmt.Printf("The size of the type of r2 (a struct) is:%d bytes\n", typ.Size())

	//23. Empty structs
	//Named
	type empty struct{}
	//Anonymous
	empty1 := struct{}{}
	fmt.Println("Empty Struct", empty1)

	//24. Interfaces - Single and Multiple
	myRect := rect1{length: 10, breadth: 20}
	myCircle := circle{radius: 30}
	fmt.Println("Rectangle via Single Interface")
	printAreaAndPerimeter(myRect)
	fmt.Println("Circle via Multiple Interfaces")
	printAreaAndPerimeter(myCircle)
	printDiameter(myCircle)

	//25. Type assertion using Interfaces
	fmt.Println("Type assertion: Shape is a rectangle or circle?")
	identifyShape(myRect)
	identifyShape(myCircle)

	//26. Type switch with Interfaces
	fmt.Println("Type switch with Interfaces:")
	determineType(myRect)
	determineType(myCircle)

	//27. Error interface
	fmt.Println("Simulate Error interface with custom error:")
	err := generateError()
	if err != nil {
		fmt.Println(err.Error())
	}

	//28. Using the errors package
	fmt.Println("Demo creating a new error using the errors package:")
	err = genErrFromPkg()
	if err != nil {
		fmt.Println(err.Error())
	}

	//29. Demo Pass through Error
	fmt.Printf("Demo of pass through error:")
	err = demoPassThroughError()
	if err != nil {
		fmt.Println(err.Error())
	}

	//30. Panic and recover
	fmt.Println("Panic and recover!")
	panic := false
	panicAndRecover(panic)

	//31. Loops in Go
	loopsInGo()

} //end of main()

// Multiple args, multiple return values, return nil error,
// argument types declared only once for similar types
func add(x, y int) (int, error) {
	x++
	return (x + y), nil
}

// Naked return
func getLocation() (x, y float64) {
	x = 38.7119882
	y = -9.3140923
	return
}

// Early Return
func earlyReturn(age int) string {
	if age < 13 {
		return "Young!"
	}
	if age < 22 {
		return "Young Adult."
	}
	if age < 45 {
		return "Middle aged adult"
	}
	if age < 67 {
		return "about to retire."
	}
	if age < 85 {
		return "Retiree."
	}
	return "Human."
}

// Functions as values
func arithmetic(a, b, c int, operation func(int, int) int) int {
	fOper := operation(a, b)
	result := operation(fOper, c)
	return result
}

func mul(x, y int) int {
	return x * y
}

func sub(f, g int) int {
	return f - g
}

// Defer functions
func completeAction(a int) {
	defer func() {
		fmt.Println("Target achieved!")
	}()

	for i := range a { //mordenized using range instead of incrementing and checking i
		fmt.Printf("Task %d done. ", i)
		if i == 5 {
			return
		}
	}
}

// Closure
func baseline() func(string) string {
	doc := ""
	return func(phrase string) string {
		doc += phrase + ". "
		return doc
	}
}

// Receiver of struct
type rect struct {
	length  int64
	breadth int64
}

func (r rect) area() int64 {
	return r.length * r.breadth
}

// Interfaces
type shape interface {
	area() float64 //Unnamed interface parameters
	perimeter() float64
}

// Interfaces: Define function with shape as args
func printAreaAndPerimeter(s shape) {
	fmt.Println("Area of the shape is:", s.area())
	fmt.Println("Perimeter of the shape is:", s.perimeter())
}

type rect1 struct {
	length  int
	breadth int
}

// Interfaces: rect1 implements interface shape
func (r rect1) area() float64 {
	return float64(r.length * r.breadth) //l*b
}
func (r rect1) perimeter() float64 {
	return float64(2 * (r.length + r.breadth)) //2(ab)
}

type circle struct {
	radius float64
}

// Interfaces: circle implements the interface shape
const pi = 3.14159

func (c circle) area() float64 {
	return pi * math.Pow(c.radius, 2) //Pir^2
}
func (c circle) perimeter() float64 {
	return 2 * pi * c.radius
}

// Multiple interfaces
type shape3d interface {
	diameter() (diameter float64) //Named interface parameters
}

// Multiple interfaces: Define function with shape3d as args
func printDiameter(s3 shape3d) {
	fmt.Println("The diameter of the 3d shape is:", s3.diameter())
}

// Multiple interfaces:circle implements diameter interface
func (c circle) diameter() float64 {
	return 2 * c.radius
}

// Type Assertion with Interfaces
func identifyShape(s shape) {
	r, ok := s.(rect1) //assert if s is a rectangle
	if ok {
		fmt.Println("The shape is a rectangle with area:", r.area())
		return
	}
	c, ok := s.(circle) //assert if s is a circle
	if ok {
		fmt.Println("The shape is a circle with perimeter:", c.perimeter())
		return
	}
}

// Type switch with interfaces
func determineType(s shape) {
	switch typ := s.(type) {
	case rect1:
		fmt.Printf("Type is %T\n", typ)
	case circle:
		fmt.Printf("Type is %T\n", typ)
	default:
		fmt.Println("Type is unknown:", typ)
	}
}

/*
Alternative type switch

	func determineType(s shape) {
		switch typ := reflect.TypeOf(s).String(); typ {
		case "main.rect1":
			fmt.Printf("Type is %s\n", typ)
		case "main.circle":
			fmt.Printf("Type is %s\n", typ)
		default:
			fmt.Println("Type is unknown", typ)
		}
	}
*/

// Implement the Error interface
func (r rect1) Error() string {
	return fmt.Sprintf("Error for rectangle with length:%d and breadth:%d", r.length, r.breadth)
}

func generateError() error {
	return rect1{length: 1000, breadth: 2000}
}

// Use the Error standard Package
func genErrFromPkg() error {
	return errors.New("generated a new error using the Errors package")
}

// Demo a pass through error
func demoPassThroughError() error {
	err := genErrFromPkg()
	if err != nil {
		pkgErr := fmt.Errorf("pass through error caught generated error using errors pkg:%w", err)
		return pkgErr
	}
	return nil
}

// Demo Panic and recover based on args
func panicAndRecover(panic bool) {
	defer func() {
		r := recover()
		if r != nil {
			fmt.Println("I have now recovered from the panic:", r)
		}
	}()
	if panic {
		createPanic()
	}
}
func createPanic() {
	panic("i am solely here to create panic!")
}

// Loops in Go.
func loopsInGo() {
	count := 3
	fmt.Printf("For loop with init, condition, increment:")
	for i := 0; i < count; i++ { //i scoped to for block
		fmt.Printf("%d;", i)
	}
	fmt.Printf("\nFor loop with init, no condition, increment:")
	for i := 0; ; i++ {
		fmt.Printf("%d;", i)
		if i > count { //prevent endless loop
			break
		}
	}
	fmt.Printf("\nFor loop with condition only like while loop:")
	j := 0 //unnecessary increase of scope using this format
	for j < count {
		fmt.Printf("%d;", j)
		j++
	}
	fmt.Printf("\nFor loop with range over counter:")
	for i := range count { //i scoped to for block
		fmt.Printf("%d;", i)
	}
}
