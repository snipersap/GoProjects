package main

import "fmt"

type shape interface {
	area() float64
	name() string
}
type triangle struct {
	height float64
	base   float64
}
type square struct {
	sideLength float64
}

func main() {

	s := square{109}
	t := triangle{height: 89, base: 96}
	print(s)
	print(t)
}

// Returns area of a triangle
func (t triangle) area() float64 {
	return t.base * t.height / 2
}

// Returns area of a square
func (s square) area() float64 {
	return s.sideLength * s.sideLength
}

// Return the name of the shape
func (t triangle) name() string {
	return "Triangle"
}

func (s square) name() string {
	return "Square"
}

// Print the area of an object
func print(sh shape) {
	fmt.Println("Area of", sh.name(), "is", sh.area())
}
