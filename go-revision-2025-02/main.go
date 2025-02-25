package main

import (
	"fmt"
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

	//19. Structs - Nested, Named and anonymous
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
		brand     string
		color     string
		model     string
		doors     int
		frontSeat seat //nested
		rearSeat  seat //nested
		wheel     struct {
			pressure float64
			kind     string
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
			pressure float64
			kind     string
		}{pressure: 3.4, kind: "M+S"}, //direct
		light: light{kind: "LED", power: "3000 Lumens", motionSensor: true}, //direct
	}

	myCar.rearSeat.upholstery = "polymer"
	myCar.rearSeat.color = "#zzzhhr"
	//Usage
	fmt.Println("My car configuration via nested, named and anonymous struct is:", myCar)

}

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
