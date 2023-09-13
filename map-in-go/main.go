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

}

func printI(ism map[int]string) {
	fmt.Println(ism)
}

func printS(ssm map[string]string) {
	fmt.Println(ssm)
}
