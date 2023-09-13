package main

import "fmt"

func main() {

	// Map storing the hex codes of colors
	colorCodes := map[string]string{
		"red":   "#ff0000",
		"green": "#008000",
		"blue":  "#0000ff",
	}
	fmt.Println(colorCodes)
	fmt.Printf("Color codes formatted:%s\n", colorCodes)

	// Map of bus routes
	var busRoutes map[int]string
	busRoutes = map[int]string{94: "Hochschule"}
	busRoutes[97] = "Londoner Ring"
	busRoutes[77] = "LU Rathaus"
	printI(busRoutes)

	//Map of train routes
	trainRoutes := make(map[string]string, 4)
	trainRoutes = map[string]string{
		"ICE5":    "Frankfurt to Basel",
		"ICE5685": "Frankfurt to Paris Est.",
		"EC52":    "Mannheim to Luzern",
		"IC90":    "Freiburg to Berlin",
	}
	printS(trainRoutes)
	delete(trainRoutes, "ICE5685") //Delete a key
	printS(trainRoutes)

	// Print color codes with loop
	fmt.Println("Print color codes (old)")
	printColors(colorCodes)

	//Update maps without pointers (maps are reference types)
	fmt.Println("Print color codes (new)")
	updateColors(colorCodes)
	printColors(colorCodes)

}

// Update color codes without using pointer
func updateColors(c map[string]string) {
	c["turquiose"] = "#00ffef"
}

// Print the colors and its hex code
func printColors(c map[string]string) {
	for color, hex := range c {
		fmt.Println("The hex code for the color", color, "is", hex)
	}
}

// Print a map of type [int]string
func printI(ism map[int]string) {
	fmt.Println(ism)
}

// Print a map of type [string]string
func printS(ssm map[string]string) {
	fmt.Println(ssm)
}
